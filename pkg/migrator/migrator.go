package migrator

import (
    "database/sql"
    "errors"
    "fmt"
    "os"
    "log"
    "path/filepath"
    "sort"
    "strings"
    "time"
)

type Migrator struct {
    db        *sql.DB
    serviceID string
    path      string
}

// NewMigrator создаёт новый экземпляр мигратора
func NewMigrator(db *sql.DB, serviceID, path string) *Migrator {
    return &Migrator{
        db:        db,
        serviceID: serviceID,
        path:      path,
    }
}

// InitSchema инициализирует таблицы для управления миграциями
func (m *Migrator) InitSchema() error {
    _, err := m.db.Exec(`
        CREATE TABLE IF NOT EXISTS migrations_lock (
            id SERIAL PRIMARY KEY,
            locked BOOLEAN DEFAULT FALSE,
            locked_at TIMESTAMPTZ DEFAULT NULL,
            locked_by TEXT DEFAULT NULL
        );
    `)
    if err != nil {
        return err
    }
    _, err = m.db.Exec(`
        CREATE TABLE IF NOT EXISTS migrations (
            id SERIAL PRIMARY KEY,
            version TEXT NOT NULL,
            applied_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            UNIQUE (version)
        );
    `)
    return err
}

// getPendingMigrations находит файлы миграций, которые ещё не применены
func (m *Migrator) getPendingMigrations() ([]string, error) {
    files, err := os.ReadDir(m.path)
    if err != nil {
        return nil, err
    }

    var migrations []string
    for _, file := range files {
        if strings.HasSuffix(file.Name(), ".sql") {
            migrations = append(migrations, file.Name())
        }
    }

    sort.Strings(migrations) // Сортировка по имени, чтобы соблюдать порядок
    return migrations, nil
}

// applyFileMigration выполняет миграцию из файла
func (m *Migrator) applyFileMigration(filename string) error {
    content, err := os.ReadFile(filepath.Join(m.path, filename))
    if err != nil {
        return err
    }

    _, err = m.db.Exec(string(content))
    if err != nil {
        return err
    }

    _, err = m.db.Exec(`INSERT INTO migrations (version) VALUES ($1)`, filename)
    return err
}

// Lock устанавливает блокировку, чтобы предотвратить параллельное выполнение миграций
func (m *Migrator) Lock() error {
    tx, err := m.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    var locked bool
    err = tx.QueryRow(`SELECT locked FROM migrations_lock WHERE id=1 FOR UPDATE`).Scan(&locked)
    if err != nil {
        // Если записи нет, создаём новую
        if errors.Is(err, sql.ErrNoRows) {
            _, err = tx.Exec(`INSERT INTO migrations_lock (id, locked, locked_at, locked_by) VALUES (1, TRUE, $1, $2)`,
                time.Now(), m.serviceID)
            if err != nil {
                return err
            }
        } else {
            return err
        }
    }

    if locked {
        return errors.New("migrations are already in progress by another service")
    }

    // Устанавливаем блокировку
    _, err = tx.Exec(`UPDATE migrations_lock SET locked=TRUE, locked_at=$1, locked_by=$2 WHERE id=1`,
        time.Now(), m.serviceID)
    if err != nil {
        return err
    }

    return tx.Commit()
}

// Unlock снимает блокировку после завершения миграции
func (m *Migrator) Unlock() error {
    _, err := m.db.Exec(`UPDATE migrations_lock SET locked=FALSE, locked_at=NULL, locked_by=NULL WHERE id=1`)
    return err
}

// ApplyMigrations выполняет миграции, которых нет в базе
func (m *Migrator) ApplyMigrations() error {
    if err := m.Lock(); err != nil {
        return fmt.Errorf("unable to acquire lock: %w", err)
    }
    defer m.Unlock()

    pendingMigrations, err := m.getPendingMigrations()
    if err != nil {
        return err
    }

    for _, migration := range pendingMigrations {
        if !m.isMigrationApplied(migration) {
            log.Printf("Applying migration: %s\n", migration)
            if err := m.applyFileMigration(migration); err != nil {
                return fmt.Errorf("failed to apply migration %s: %w", migration, err)
            }
        }
    }

    return nil
}

// isMigrationApplied проверяет, была ли миграция выполнена ранее
func (m *Migrator) isMigrationApplied(version string) bool {
    var exists bool
    err := m.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM migrations WHERE version = $1)`, version).Scan(&exists)
    if err != nil {
        log.Printf("Failed to check migration version: %v", err)
    }
    return exists
}

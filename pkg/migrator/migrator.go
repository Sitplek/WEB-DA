=package migrator

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Migrator struct {
	db        *sqlx.DB
	serviceID string
	path      string
}

// NewMigrator создает новый экземпляр мигратора
func NewMigrator(db *sqlx.DB, serviceID, path string) *Migrator {
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

// ApplyMigrations выполняет миграции, которых нет в базе
func (m *Migrator) ApplyMigrations() error {
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

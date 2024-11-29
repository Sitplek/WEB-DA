package main

import (
    "database/sql"
    "log"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
	_"github.com/lib/pq"
    "gorm.io/gorm"

    "WEB-DA/internal/handlers"
    "WEB-DA/internal/service"
    "WEB-DA/internal/storage"
    "WEB-DA/pkg/migrator"
)

func main() {
    // Подключение к базе данных
    // dsn := "host=localhost user=test_user password=test_password dbname=test_db port=5432 sslmode=disable"
    sqlDB, err := sql.Open("postgres", "postgres://test_user:test_password@postgres:5432/test_db?sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }
    defer sqlDB.Close()

    // Применение миграций
    mig := migrator.NewMigrator(sqlDB, "organization_service", "./migrations")
    if err := mig.InitSchema(); err != nil {
        log.Fatalf("Failed to init schema: %v", err)
    }
    if err := mig.ApplyMigrations(); err != nil {
        log.Fatalf("Failed to apply migrations: %v", err)
    }

    // Инициализация GORM
    gormDB, err := gorm.Open(postgres.New(postgres.Config{
        Conn: sqlDB,
    }), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to initialize GORM: %v", err)
    }

    // Регистрация зависимостей
    storage := storage.NewOrganizationStorage(gormDB)
    service := service.NewOrganizationService(storage)
    handler := handlers.NewOrganizationHandler(service)

    // Создание роутера
    r := gin.Default()
    r.GET("/organization/tree", handler.GetOrganizationTree)
    r.GET("/employees", handler.GetFilteredEmployees)

    // Запуск сервиса
    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

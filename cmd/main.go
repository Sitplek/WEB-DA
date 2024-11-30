package main

import (
    "log"
    "github.com/gin-gonic/gin"
	_"github.com/lib/pq"
    "github.com/jmoiron/sqlx"

    "WEB-DA/internal/handler"
    "WEB-DA/internal/service"
    "WEB-DA/internal/storage"
    "WEB-DA/pkg/migrator"
)

func main() {
	// Подключение к базе данных
	db, err := sqlx.Connect("postgres", "postgresql://test_user:test_password@postgres:5432/test_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

    // Применение миграций
    mig := migrator.NewMigrator(db, "organization_service", "./migrations")
    if err := mig.InitSchema(); err != nil {
        log.Fatalf("Failed to init schema: %v", err)
    }
    if err := mig.ApplyMigrations(); err != nil {
        log.Fatalf("Failed to apply migrations: %v", err)
    }

    // Регистрация зависимостей
	storageLayer := storage.NewStorage(db)
	serviceLayer := service.NewService(storageLayer)
	handler := handler.NewHandler(serviceLayer)

    // Создание роутера
	r := gin.Default()
	r.GET("/employees/hierarchy", handler.GetEmployeeHierarchy)
	r.GET("/employees", handler.GetAllEmployees)

    // Запуск сервиса
    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

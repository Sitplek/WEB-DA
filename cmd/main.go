// Package main
// @title Employee Management API
// @version 1.0
// @description API для управления сотрудниками и их иерархией.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://example.com/contact
// @contact.email support@example.com
// @host localhost:8080
// @BasePath /
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"

	"WEB-DA/internal/handler"
	"WEB-DA/internal/middleware"
	"WEB-DA/internal/service"
	"WEB-DA/internal/storage"
	"WEB-DA/pkg/migrator"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Подключение CORS middleware
	r.Use(middleware.CORS())

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Роуты приложения
	r.GET("/employees/hierarchy", handler.GetEmployeeHierarchy)
	r.GET("/employees", handler.GetAllEmployees)

	// Запуск сервиса
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

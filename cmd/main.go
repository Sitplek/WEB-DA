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
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	"WEB-DA/internal/handler"
	"WEB-DA/internal/middleware"
	"WEB-DA/internal/service"
	"WEB-DA/internal/storage"
	"WEB-DA/pkg/migrator"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Подключение к PostgreSQL
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Подключение к Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Применение миграций
	mig := migrator.NewMigrator(db, "organization_service", "./migrations")
	if err := mig.InitSchema(); err != nil {
		log.Fatalf("Failed to init schema: %v", err)
	}
	if err := mig.ApplyMigrations(); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	// Инициализация слоев
	storageLayer := storage.NewStorage(db, rdb)
	serviceLayer := service.NewService(storageLayer)
	handler := handler.NewHandler(serviceLayer)

	// Настройка роутера
	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/employees", handler.GetAllEmployees)
	r.GET("/employees/hierarchy", handler.GetEmployeeHierarchy)

	// Запуск сервера
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

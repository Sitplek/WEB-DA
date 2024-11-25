package main

import (
	"database/sql"
	"log"
	"os"
	"WEB-DA/internal/handlers"
	"WEB-DA/internal/service"
	"WEB-DA/internal/storage"
	"WEB-DA/pkg/migrator"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL драйвер
)

func main() {
	// Загружаем переменные окружения из .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, falling back to environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	serviceID := os.Getenv("SERVICE_ID")
	port := os.Getenv("PORT")
	if dbURL == "" || serviceID == "" || port == "" {
		log.Fatal("Missing required environment variables: DATABASE_URL, SERVICE_ID, or PORT")
	}


	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Создание мигратора
	m := migrator.NewMigrator(db, "WEB-DA", "./migrations")

	// Инициализация схемы мигратора
	if err := m.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize migration schema: %v", err)
	}

	// Применение миграций
	if err := m.ApplyMigrations(); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully. Starting the service...")

	storage := storage.NewStorage(db)
	service := service.NewService(storage)
	handlers := handlers.NewHandlers(service)

	router := gin.Default()

	// Запросы
	router.GET("/organization-tree", handlers.GetOrganizationTree)
	router.GET("/employee-search", handlers.EmployeeSearch)
	router.GET("/employee-details/:id", handlers.GetEmployeeDetails)

	// Мутации
	router.PUT("/update-employee", handlers.UpdateEmployee)
	router.POST("/add-department", handlers.AddDepartment)

	log.Println("Server is running on :8080")
	router.Run(":8080")
}

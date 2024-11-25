package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	ServiceID   string
	Port        string
	Postgres    struct {
		User     string
		Password string
		DBName   string
	}
}

// LoadConfig загружает конфигурацию из файла .env и переменных окружения
func LoadConfig() *Config {
	// Загружаем .env файл, если он существует
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Создаём экземпляр конфигурации
	cfg := &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		ServiceID:   getEnv("SERVICE_ID", "default-service"),
		Port:        getEnv("PORT", "8080"),
	}

	// Настройка PostgreSQL
	cfg.Postgres.User = getEnv("POSTGRES_USER", "postgres")
	cfg.Postgres.Password = getEnv("POSTGRES_PASSWORD", "postgres")
	cfg.Postgres.DBName = getEnv("POSTGRES_DB", "my_database")

	return cfg
}

// getEnv возвращает значение переменной окружения с поддержкой значений по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

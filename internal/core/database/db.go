package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	if err := godotenv.Load(".env.local"); err == nil {
		log.Println("✅ Loaded config from .env.local")
	}

	if err := godotenv.Load(".env"); err == nil {
		log.Println("✅ Loaded config from .env")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	log.Println("✅ Database connected successfully!")

	return db
}

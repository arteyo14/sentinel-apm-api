package main

import (
	"log"
	"os"
	"sentinel-apm-api/internal/core/database"
	"sentinel-apm-api/internal/core/migrate"
	"sentinel-apm-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()

	if db == nil {
		log.Fatal("failed to connect to database")
	}

	migrate.Migrate(db)

	router := gin.Default()
	routes.SetupRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on port %s", port)
	router.Run(":" + port)
}

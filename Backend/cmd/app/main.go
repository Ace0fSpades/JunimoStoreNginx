package main

import (
	"log"
	"os"

	docs "uniStore/Backend/api/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"uniStore/Backend/internal/infrastructure/database"
	"uniStore/Backend/internal/interfaces/api"
	"uniStore/Backend/internal/utils"
)

// initConfig loads environment variables from .env file
func initConfig() error {
	// Load .env file but don't return error if it doesn't exist
	// This allows environment variables to be set directly in docker-compose
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found or could not be loaded: %v", err)
		log.Println("Continuing with environment variables set in the system")
		return nil
	}
	return nil
}

// @title		Game Store
// @version		1.0
// @description	REST-API for game store

// @host		localhost
// @basePath	/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description JWT Authorization header using Bearer scheme. Example: "Bearer {token}"
func main() {
	// Setup logging
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)

	// Initialize configuration
	if err := initConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	if err := db.CheckAdminAndRoles(); err != nil {
		log.Fatalf("Failed to check admin and roles: %v", err)
	}

	// Initialize router and services
	router := gin.Default()
	server := api.NewServer(db, router)
	server.SetupRoutes()

	// Configure Swagger
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	docs.SwaggerInfo.Host = "localhost:" + port

	if !utils.IsProd() {
		log.Printf("Swagger UI is available at: http://127.0.0.1:%s/swagger/index.html\n", port)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server on port %s: %v", port, err)
	}
}

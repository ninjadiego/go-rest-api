package main

import (
  	"log"
  	"os"

  	"github.com/gin-gonic/gin"
  	"github.com/joho/godotenv"
  	"github.com/ninjadiego/go-rest-api/internal/handlers"
  	"github.com/ninjadiego/go-rest-api/internal/middleware"
  	"github.com/ninjadiego/go-rest-api/config"
  )

func main() {
  	// Load environment variables
  	if err := godotenv.Load(); err != nil {
      		log.Println("No .env file found, using environment variables")
      	}

  	// Initialize database
  	db, err := config.InitDB()
  	if err != nil {
      		log.Fatalf("Failed to connect to database: %v", err)
      	}

  	// Auto-migrate models
  	config.AutoMigrate(db)

  	// Initialize Gin router
  	r := gin.Default()

  	// Initialize handlers
  	authHandler := handlers.NewAuthHandler(db)
  	userHandler := handlers.NewUserHandler(db)

  	// Public routes
  	api := r.Group("/api/v1")
  	{
      		auth := api.Group("/auth")
      		{
            			auth.POST("/register", authHandler.Register)
            			auth.POST("/login", authHandler.Login)
            		}
      	}

  	// Protected routes
  	protected := api.Group("/users")
  	protected.Use(middleware.JWTMiddleware())
  	{
      		protected.GET("/", userHandler.GetAll)
      		protected.GET("/:id", userHandler.GetByID)
      		protected.PUT("/:id", userHandler.Update)
      		protected.DELETE("/:id", userHandler.Delete)
      	}

  	// Start server
  	port := os.Getenv("PORT")
  	if port == "" {
      		port = "8080"
      	}

  	log.Printf("Server running on port %s", port)
  	if err := r.Run(":" + port); err != nil {
      		log.Fatalf("Failed to start server: %v", err)
      	}
  }

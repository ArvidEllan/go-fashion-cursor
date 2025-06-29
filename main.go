package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"virtual-tryon/config"
	"virtual-tryon/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database connection
	config.InitializeDatabase()

	// Initialize router
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Initialize routes
	initializeRoutes(router)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// initializeRoutes sets up all the routes for the application
func initializeRoutes(router *gin.Engine) {
	// API group
	api := router.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", handleRegister)
			auth.POST("/login", handleLogin)
		}

		// User routes
		user := api.Group("/user")
		{
			user.GET("/profile", middleware.AuthMiddleware(), handleGetProfile)
			user.PUT("/profile", middleware.AuthMiddleware(), handleUpdateProfile)
		}

		// Try-on routes
		tryon := api.Group("/try-on")
		{
			tryon.POST("/upload", middleware.AuthMiddleware(), handleUploadPhoto)
			tryon.POST("/process", middleware.AuthMiddleware(), handleProcessTryOn)
			tryon.GET("/history", middleware.AuthMiddleware(), handleGetTryOnHistory)
			tryon.DELETE("/history/:id", middleware.AuthMiddleware(), handleDeleteTryOnHistory)
		}

		// Product routes
		products := api.Group("/products")
		{
			products.GET("", handleGetProducts)
			products.GET("/:id", handleGetProduct)
		}

		// Cart routes
		cart := api.Group("/cart")
		{
			cart.POST("", middleware.AuthMiddleware(), handleAddToCart)
			cart.GET("", middleware.AuthMiddleware(), handleGetCart)
		}
	}
}

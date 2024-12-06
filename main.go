package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(".env.example"); err != nil {
		log.Printf("Warning: .env file not found or error loading. Using default values. Error: %v", err)
	}
}

func main() {
	// Initialize router
	r := gin.Default()
	
	// Initialize services
	promoService := NewPromoService()
	productService := NewProductService()
	orderService := NewOrderService(productService)
	
	// Initialize handlers
	productHandler := NewProductHandler(productService)
	orderHandler := NewOrderHandler(orderService, promoService)
	
	// Setup middleware
	r.Use(ErrorMiddleware())
	
	// Setup routes
	api := r.Group("/api")
	{
		// Product routes
		api.GET("/product", productHandler.ListProducts)
		api.GET("/product/:productId", productHandler.GetProduct)
		
		// Order routes
		api.POST("/order", AuthMiddleware(), orderHandler.PlaceOrder)
	}
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
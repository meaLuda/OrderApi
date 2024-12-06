package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("api_key")
		expectedKey := os.Getenv("API_KEY")
		if expectedKey == "" {
			expectedKey = "apitest" // Default value if not set in env
		}
		
		if apiKey != expectedKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API key",
			})
			return
		}
		c.Next()
	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": c.Errors,
			})
		}
	}
}
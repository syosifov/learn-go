package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Gin!",
	})
}

func VarsTest(c *gin.Context) {
	variables := []string{"APP_PORT",
		"DB_HOST",
		"DB_USER",
		"DB_PASSWORD",
		"MY_MESSAGE"}
	varsMap := gin.H{}

	// Retrieve and print the values
	for _, key := range variables {
		value := os.Getenv(key)
		if value == "" {
			log.Printf("Warning: %s is not set\n", key)
		} else {
			fmt.Printf("%s: %s\n", key, value)
			varsMap[key] = value
		}
	}
	c.JSON(200, varsMap)
}

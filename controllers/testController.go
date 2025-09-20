package controllers

import (
	"fmt"
	"log"
	tlg "main/telegram"
	"os"
	"time"

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

func TSendMessage(c *gin.Context) {
	type RequestBody struct {
		ChatID int64  `json:"chat_id"`
		Text   string `json:"text"`
	}

	var reqBody RequestBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if reqBody.ChatID == 0 || reqBody.Text == "" {
		c.JSON(400, gin.H{"error": "chat_id and text are required"})
		return
	}

	// Here you would typically send the message using your Telegram bot logic.
	// For demonstration purposes, we'll just log the message.
	log.Printf("Sending message to chat_id %d: %s", reqBody.ChatID, reqBody.Text)

	for i := 0; i < 3; i++ {
		tlg.SendMessage(reqBody.ChatID, reqBody.Text)
		if i == 2 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	c.JSON(200, gin.H{"status": "Message sent"})
}

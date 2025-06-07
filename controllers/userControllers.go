package controllers

import (

	// "main/auth"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List users",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create user",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get user",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update user",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete user",
	})
}

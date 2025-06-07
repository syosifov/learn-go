package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"main/utils" // Replace with your actual project path

	"github.com/gin-gonic/gin"
)

// ResetPasswordRequest represents the request for a password reset
type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// ForgotPassword handles password reset requests
func ForgotPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if user exists
	// This is just a placeholder - implement your user lookup logic here
	// user, exists := findUserByEmail(req.Email)
	// if !exists {
	//     // We still return success to avoid leaking user information
	//     c.JSON(http.StatusOK, gin.H{"message": "If your email is registered, you will receive reset instructions"})
	//     return
	// }

	// Generate a random token
	token := generateResetToken()

	// Store the token in your database with an expiration time
	// storeResetToken(user.ID, token, time.Now().Add(30*time.Minute))

	// Send the reset email
	emailConfig := utils.LoadEmailConfig()
	err := utils.SendPasswordResetEmail(req.Email, token, emailConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send reset email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "If your email is registered, you will receive reset instructions"})
}

// Generate a random token for password reset
func generateResetToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

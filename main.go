package main

import (
	"cmp"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize()

	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Initialize the CORS
	// initCors(r)
	// registerRoutes(r)

	r.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "0.0.1 / 2025-05-04"})
	})

	port := cmp.Or(os.Getenv("PORT"), "8080")

	// logger.Info("Server starting", slog.String("port", port))
	log.Printf("Server starting at port %s \n", port)

	// go telegram.RunTelegram()
	r.Run((":" + port))
}

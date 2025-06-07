package main

import (
	"fmt"
	"log"
	"os"

	"main/database"
	"main/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initialize() {
	// Initialize the application
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, proceeding with environment variables.")
	}
	variables := []string{"APP_PORT",
		"PORT",
		"MY_MESSAGE"}

	// Retrieve and print the values
	for _, key := range variables {
		value := os.Getenv(key)
		if value == "" {
			log.Printf("Warning: %s is not set\n", key)
		} else {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	database.DBInit()
	// telegram.RunTelegram()

}

func initCors(r *gin.Engine) {
	// Initialize the CORS
	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"origin", "content-type", "authorization"}
	config.AllowCredentials = true

	r.Use(cors.New(config))
	// router.Run()
}

func registerRoutes(r *gin.Engine) {

	// routes.RegisterUserRoutes(r)    // /api/users
	routes.RegisterTestRoutes(r)  // /t
	routes.RegisterAdminRoutes(r) // /api/admin
	// routes.RegisterManagerRoutes(r) // /api/manager
	// routes.RegisterCahier(r)        // /api/cahier
}

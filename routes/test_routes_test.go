package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return router
}

// TestPingRoute verifies that the GET /ping endpoint returns a 200 status code
// and "pong" as the response body. It uses the Arrange-Act-Assert pattern to:
// 1. Set up a test router instance and prepare an HTTP GET request to /ping
// 2. Execute the request against the router
// 3. Validate that the response status is 200 and the body contains "pong"
func TestPingRoute(t *testing.T) {
	// 1. Arrange

	router := setupRouter() // Function that initializes your Gin engine
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	// 2. Act
	router.ServeHTTP(w, req)

	// 3. Assert
	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	if w.Body.String() != "pong" {
		t.Errorf("Expected 'pong', got %s", w.Body.String())
	}
}

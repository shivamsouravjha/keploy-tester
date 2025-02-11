package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRecoveryMiddleware_PanicRecovery(t *testing.T) {
	// Create a new Gin router with the middleware
	router := gin.New()
	router.Use(RecoveryMiddleware())

	// Define a handler that will panic
	router.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	// Perform a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/panic", nil)
	router.ServeHTTP(w, req)

	// Check that the response status is 500 Internal Server Error
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}

	// Check that the response body is as expected
	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	if response["error"] != "Internal server error. Please try again later." {
		t.Errorf("Expected error message %q, got %q", "Internal server error. Please try again later.", response["error"])
	}
}

func TestRecoveryMiddleware_NoPanic(t *testing.T) {
	// Create a new Gin router with the middleware
	router := gin.New()
	router.Use(RecoveryMiddleware())

	// Define a handler that does not panic
	router.GET("/no_panic", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Perform a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/no_panic", nil)
	router.ServeHTTP(w, req)

	// Check that the response status is 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check that the response body is as expected
	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	if response["message"] != "success" {
		t.Errorf("Expected message %q, got %q", "success", response["message"])
	}
}

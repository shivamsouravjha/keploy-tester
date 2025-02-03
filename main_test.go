package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCORSMiddleware_WithOriginHeader(t *testing.T) {
	// Create a test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set the Origin header in the request
	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Origin", "http://example.com")

	// Apply the middleware
	middleware := CORSMiddleware()
	middleware(c)

	// Check the headers
	if w.Header().Get("Access-Control-Allow-Origin") != "http://example.com" {
		t.Errorf("Expected Access-Control-Allow-Origin to be 'http://example.com', got '%s'", w.Header().Get("Access-Control-Allow-Origin"))
	}
	if w.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Errorf("Expected Access-Control-Allow-Credentials to be 'true', got '%s'", w.Header().Get("Access-Control-Allow-Credentials"))
	}
	if w.Header().Get("Access-Control-Allow-Headers") == "" {
		t.Errorf("Expected Access-Control-Allow-Headers to be set, but it was empty")
	}
	if w.Header().Get("Access-Control-Allow-Methods") == "" {
		t.Errorf("Expected Access-Control-Allow-Methods to be set, but it was empty")
	}
}

func TestCORSMiddleware_OptionsMethod(t *testing.T) {
	// Create a test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set the HTTP method to OPTIONS
	c.Request = httptest.NewRequest("OPTIONS", "/", nil)

	// Apply the middleware
	middleware := CORSMiddleware()
	middleware(c)

	// Check the response status code
	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status code 204, got %d", w.Code)
	}
}

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Test generated using Keploy
func TestJWTAuthMiddleware_InvalidToken_Returns401(t *testing.T) {
	// Create a test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set an invalid Authorization header
	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Authorization", "Bearer invalid_token")

	// Call the middleware
	middleware := JWTAuthMiddleware()
	middleware(c)

	// Assert that the response status code is 401
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %d", w.Code)
	}

	// Assert the response body contains the correct error message
	expectedBody := `{"error":"Invalid token"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}

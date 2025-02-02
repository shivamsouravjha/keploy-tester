package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProtectedPurgeEventsRoute_Unauthorized(t *testing.T) {
	r := gin.Default()
	Routes(r)

	req, _ := http.NewRequest("DELETE", "/api/events/purge", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d for unauthorized access, but got %d", http.StatusUnauthorized, w.Code)
	}
}

package postgres

// Test generated using Keploy
import (
    "testing"
    "gorm.io/gorm"
)

func TestGetDB_ReturnsInitializedDB(t *testing.T) {
    // Mock the DB initialization
    DB = &gorm.DB{}
    
    // Call the function
    result := GetDB()

    // Validate the result
    if result == nil {
        t.Errorf("Expected non-nil *gorm.DB instance, got nil")
    }
}

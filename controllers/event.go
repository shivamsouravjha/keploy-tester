package controllers

import (
	"encoding/json"
	"net/http"
	"segwise/clients/postgres"
	redis_client "segwise/clients/redis"
	"segwise/models"
	"time"

	"github.com/gin-gonic/gin"
)

// ExecuteTrigger manually triggers an event
// @Summary Manually execute a trigger
// @Description Triggers an event immediately for testing
// @Tags Triggers
// @Produce json
// @Param id path string true "Trigger ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/triggers/{id}/execute [post]
func ExecuteTrigger(c *gin.Context) {

	id := c.Param("id")
	var trigger models.Trigger
	db := postgres.GetDB()

	if err := db.First(&trigger, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trigger not found"})
		return
	}

	// Log the event
	event := models.EventLog{
		TriggerID:   id,
		TriggeredAt: time.Now(),
		Payload:     trigger.Payload,
		Type:        trigger.Type,
		Status:      "active",
	}

	db.Create(&event)

	// Store in Redis for fast lookup
	eventJSON, _ := json.Marshal(event)
	redisClient := redis_client.RedisSession()

	redisClient.Set("event:"+id, eventJSON, 2*time.Hour)

	c.JSON(http.StatusOK, gin.H{"message": "Trigger executed", "event": event})

}

// GetActiveEvents retrieves recent event logs
func GetActiveEvents(c *gin.Context) {
	db := postgres.GetDB()

	var events []models.EventLog
	db.Where("status = ?", "active").Find(&events)
	c.JSON(http.StatusOK, events)

}

// GetArchivedEvents retrieves archived event logs
func GetArchivedEvents(c *gin.Context) {
	db := postgres.GetDB()

	var events []models.EventLog
	db.Where("status = ?", "archived").Find(&events)
	c.JSON(http.StatusOK, events)

}

// PurgeOldEvents deletes expired logs
func PurgeOldEvents(c *gin.Context) {
	db := postgres.GetDB()

	// Delete logs older than 48 hours
	timeThreshold := time.Now().Add(-48 * time.Hour)
	db.Where("triggered_at < ?", timeThreshold).Delete(&models.EventLog{})

	c.JSON(http.StatusOK, gin.H{"message": "Old events purged"})

}

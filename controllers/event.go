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
// @Router /triggers/{id}/execute [post]
// @Security BearerAuth
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

// @Summary Get active event logs
// @Description Fetch event logs from the last 2 hours
// @Produce  json
// @Success 200 {array} models.EventLog
// @Router /events [get]
// @Security BearerAuth
func GetActiveEvents(c *gin.Context) {
	db := postgres.GetDB()

	var events []models.EventLog
	db.Where("status = ?", "active").Find(&events)
	c.JSON(http.StatusOK, events)

}

// @Summary Get archived event logs
// @Description Fetch logs that are older than 2 hours but still within 48 hours
// @Produce  json
// @Success 200 {array} models.EventLog
// @Router /events/archived [get]
// @Security BearerAuth
func GetArchivedEvents(c *gin.Context) {
	db := postgres.GetDB()

	var events []models.EventLog
	db.Where("status = ?", "archived").Find(&events)
	c.JSON(http.StatusOK, events)

}

// @Summary Purge old event logs
// @Description Deletes logs older than 48 hours
// @Success 200 {string} string "Old events purged"
// @Router /events/purge [delete]
// @Security BearerAuth
func PurgeOldEvents(c *gin.Context) {
	db := postgres.GetDB()

	// Delete logs older than 48 hours
	timeThreshold := time.Now().Add(-48 * time.Hour)
	db.Where("triggered_at < ?", timeThreshold).Delete(&models.EventLog{})

	c.JSON(http.StatusOK, gin.H{"message": "Old events purged"})

}

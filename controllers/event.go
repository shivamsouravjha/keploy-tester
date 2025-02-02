package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"segwise/clients/postgres"
	redis_client "segwise/clients/redis"
	"segwise/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	redisClient := redis_client.RedisSession()
	cached, err := redisClient.Get("activeEvents").Result()
	if err == nil && cached != "" {
		// We have data in Redis
		var events []models.EventLog
		if jsonErr := json.Unmarshal([]byte(cached), &events); jsonErr == nil {
			c.JSON(http.StatusOK, events)
			return
		}
	}

	// Otherwise, fall back to DB
	db := postgres.GetDB()
	var events []models.EventLog
	db.Where("status = ?", "active").Find(&events)

	// Cache the results in Redis for next time
	eventBytes, _ := json.Marshal(events)
	redisClient.Set("activeEvents", eventBytes, 2*time.Hour)

	c.JSON(http.StatusOK, events)
}

// @Summary Get archived event logs
// @Description Fetch logs that are older than 2 hours but still within 48 hours
// @Produce  json
// @Success 200 {array} models.EventLog
// @Router /events/archived [get]
// @Security BearerAuth
func GetArchivedEvents(c *gin.Context) {
	redisClient := redis_client.RedisSession()
	cached, err := redisClient.Get("archivedEvents").Result()
	if err == nil && cached != "" {
		var events []models.EventLog
		if jsonErr := json.Unmarshal([]byte(cached), &events); jsonErr == nil {
			c.JSON(http.StatusOK, events)
			return
		}
	}

	db := postgres.GetDB()

	var events []models.EventLog
	db.Where("status = ?", "archived").Find(&events)
	eventBytes, _ := json.Marshal(events)
	redisClient.Set("archivedEvents", eventBytes, 2*time.Hour)

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
	redisClient := redis_client.RedisSession()
	redisClient.Del("activeEvents")
	redisClient.Del("archivedEvents")
	c.JSON(http.StatusOK, gin.H{"message": "Old events purged"})

}

type ScheduledTriggerTestRequest struct {
	Delay int `json:"delay"` // Delay in minutes
}

// TestScheduledTrigger
// @Summary Test a one-time scheduled trigger
// @Description This endpoint allows users to test a scheduled event trigger **without saving it permanently**.
// @Tags Testing API
// @Accept json
// @Produce json
// @Param request body ScheduledTriggerTestRequest true "Request body for scheduled trigger test"
// @Success 200 {object} map[string]interface{} "Trigger executed successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Router /triggers/test/scheduled [post]
// @Security BearerAuth
func TestScheduledTrigger(c *gin.Context) {
	redisClient := redis_client.RedisSession()
	var request ScheduledTriggerTestRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	go func() {
		time.Sleep(time.Duration(request.Delay) * time.Minute)

		// Simulate event execution
		event := models.EventLog{
			TriggeredAt: time.Now(),
			Type:        "scheduled",
			Status:      "test",
		}

		// Store event temporarily in Redis for quick lookup (expires in 1 hour)
		eventJSON, _ := json.Marshal(event)
		redisClient.Set("test_event", eventJSON, 1*time.Hour)

		zap.L().Info("One-time scheduled trigger executed for testing")
	}()

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Test scheduled trigger will fire in %d", request.Delay), "delay": request.Delay, "minutes": true})

}

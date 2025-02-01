package controllers

import (
	"net/http"
	"segwise/clients/postgres"
	"segwise/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateTrigger handles creating a new trigger
func CreateTrigger(c *gin.Context) {
	db := postgres.GetDB()

	var trigger models.Trigger
	if err := c.ShouldBindJSON(&trigger); err != nil {
		zap.L().Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if trigger.Type != "scheduled" && trigger.Type != "api" {
		zap.L().Error("Trigger type must be 'scheduled' or 'api'")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Trigger type must be 'scheduled' or 'api'"})
		return
	}

	if err := db.Create(&trigger).Error; err != nil {
		zap.L().Error("Failed to create trigger", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create trigger"})
		return
	}

	c.JSON(http.StatusCreated, trigger)

}

// GetTriggers retrieves all triggers
func GetTriggers(c *gin.Context) {
	var triggers []models.Trigger
	db := postgres.GetDB()
	db.Find(&triggers)
	c.JSON(http.StatusOK, triggers)

}

// GetTriggerByID retrieves a specific trigger
func GetTriggerByID(c *gin.Context) {
	var trigger models.Trigger
	id := c.Param("id")
	db := postgres.GetDB()
	if err := db.First(&trigger, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trigger not found"})
		return
	}

	c.JSON(http.StatusOK, trigger)

}

// UpdateTrigger modifies an existing trigger
func UpdateTrigger(c *gin.Context) {

	id := c.Param("id")
	var trigger models.Trigger
	db := postgres.GetDB()

	if err := db.First(&trigger, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trigger not found"})
		return
	}

	if err := c.ShouldBindJSON(&trigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&trigger)
	c.JSON(http.StatusOK, trigger)

}

// DeleteTrigger removes a trigger
func DeleteTrigger(c *gin.Context) {
	db := postgres.GetDB()

	id := c.Param("id")
	if err := db.Delete(&models.Trigger{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete trigger"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Trigger deleted"})

}

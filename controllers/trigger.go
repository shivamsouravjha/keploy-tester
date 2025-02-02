package controllers

import (
	"net/http"
	"segwise/clients/postgres"
	"segwise/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateTrigger creates a new event trigger
// @Summary Create a new trigger
// @Description Creates a scheduled or API trigger
// @Tags Triggers
// @Accept json
// @Produce json
// @Param trigger body models.Trigger true "Trigger object"
// @Success 201 {object} models.Trigger
// @Failure 400 {object} map[string]string
// @Router /triggers [post]
// @Security BearerAuth
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
// @Summary Get all triggers
// @Description Fetch all stored triggers
// @Tags Triggers
// @Produce json
// @Success 200 {array} models.Trigger
// @Router /triggers [get]
// @Security BearerAuth
func GetTriggers(c *gin.Context) {
	var triggers []models.Trigger
	db := postgres.GetDB()
	db.Find(&triggers)
	c.JSON(http.StatusOK, triggers)

}

// @Summary Get a specific trigger
// @Description Fetch a trigger by ID
// @Produce  json
// @Param id path string true "Trigger ID"
// @Success 200 {object} models.Trigger
// @Router /triggers/{id} [get]
// @Security BearerAuth
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

// @Summary Update a trigger
// @Description Modify an existing trigger
// @Accept  json
// @Produce  json
// @Param id path string true "Trigger ID"
// @Param trigger body models.Trigger true "Updated Trigger object"
// @Success 200 {object} models.Trigger
// @Router /triggers/{id} [put]
// @Security BearerAuth
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

// @Summary Delete a trigger
// @Description Removes a trigger from the system
// @Param id path string true "Trigger ID"
// @Success 200 {string} string "Trigger deleted"
// @Router /triggers/{id} [delete]
// @Security BearerAuth
func DeleteTrigger(c *gin.Context) {
	db := postgres.GetDB()

	id := c.Param("id")
	if err := db.Delete(&models.Trigger{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete trigger"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Trigger deleted"})

}

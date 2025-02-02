package helpers

import (
	"segwise/clients/postgres"
	"segwise/models"
	"time"

	cron "github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func StartEventRetentionScheduler() {
	db := postgres.GetDB()
	c := cron.New()

	// Archive events older than 2 hours (Runs every 10 minutes)
	_, err := c.AddFunc("*/10 * * * *", func() {
		zap.L().Info("Archiving old events...")
		db.Model(&models.EventLog{}).
			Where("status = ?", "active").
			Where("triggered_at < ?", time.Now().Add(-2*time.Hour)).
			Update("status", "archived")
		zap.L().Info("Archived old events.")
	})
	if err != nil {
		zap.L().Error("Failed to schedule event archival:", zap.Error(err))
	}

	// Delete events older than 48 hours (Runs every 30 minutes)
	_, err = c.AddFunc("*/30 * * * *", func() {
		zap.L().Info("Deleting expired events...")
		db.Where("status = ?", "archived").
			Where("triggered_at < ?", time.Now().Add(-48*time.Hour)).
			Delete(&models.EventLog{})
		zap.L().Info("Deleted expired events.")
	})
	if err != nil {
		zap.L().Error("Failed to schedule event deletion:", zap.Error(err))
	}

	// Start the cron scheduler
	c.Start()
	zap.L().Info("Event retention scheduler started.")
}

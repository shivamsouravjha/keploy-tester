package helpers

import (
	"encoding/json"
	"log"
	"segwise/clients/postgres"
	redis_client "segwise/clients/redis"
	"segwise/models"
	"sync"
	"time"

	"github.com/go-redis/redis"
	cron "github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Scheduler struct {
	Cron  *cron.Cron
	Jobs  map[string]cron.EntryID // Store triggerID -> Job ID
	Mutex sync.Mutex              // Prevent race conditions
	DB    *gorm.DB
	Redis *redis.Client
}

// StartScheduler initializes the scheduler and continuously checks for updates
func StartScheduler() *Scheduler {
	db := postgres.GetDB()
	redisClient := redis_client.RedisSession()
	s := &Scheduler{
		Cron:  cron.New(),
		Jobs:  make(map[string]cron.EntryID),
		DB:    db,
		Redis: redisClient,
	}

	// Load initial triggers
	s.syncTriggers()

	// Start a goroutine to refresh triggers every 1 minute
	go func() {
		for {
			time.Sleep(60 * time.Second)
			s.syncTriggers()
		}
	}()

	// Start the cron scheduler
	s.Cron.Start()
	return s
}

// Sync database triggers with running jobs
func (s *Scheduler) syncTriggers() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	var triggers []models.Trigger
	s.DB.Where("type = ?", "scheduled").Find(&triggers)

	activeTriggerIDs := make(map[string]bool)

	// Add or update triggers
	for _, trigger := range triggers {
		activeTriggerIDs[trigger.ID] = true

		// If trigger is new, add it
		if _, exists := s.Jobs[trigger.ID]; !exists {
			jobID, err := s.Cron.AddFunc(trigger.Schedule, func() {
				s.executeScheduledTrigger(trigger)
			})
			if err != nil {
				log.Println("Failed to schedule trigger:", err)
				continue
			}
			s.Jobs[trigger.ID] = jobID
			log.Printf("Scheduled new trigger: %s (%s)", trigger.ID, trigger.Schedule)
		}
	}

	// Remove deleted triggers
	for triggerID, jobID := range s.Jobs {
		if !activeTriggerIDs[triggerID] {
			s.Cron.Remove(jobID)
			delete(s.Jobs, triggerID)
			log.Printf("Removed trigger: %s", triggerID)
		}
	}
}

// Execute the scheduled trigger
func (s *Scheduler) executeScheduledTrigger(trigger models.Trigger) {
	event := models.EventLog{
		TriggerID:   trigger.ID,
		TriggeredAt: time.Now(),
		Payload:     trigger.Payload,
		Type:        "scheduled",
		Status:      "active",
	}

	s.DB.Create(&event)

	// Cache in Redis for fast lookup
	eventJSON, _ := json.Marshal(event)
	s.Redis.Set("event:"+trigger.ID, eventJSON, 2*time.Hour)

	log.Printf("Executed scheduled trigger: %s", trigger.ID)
}

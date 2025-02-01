package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"segwise/clients/postgres"
	"segwise/models"
	"time"

	redis_client "segwise/clients/redis"

	"github.com/go-redis/redis"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func StartScheduler() {

	cron := cron.New()
	db := postgres.GetDB()
	redisClient := redis_client.RedisSession()

	// Fetch from DB, but use a new slice!
	var storedTriggers []models.Trigger
	err := db.Where("type = ?", "scheduled").Find(&storedTriggers).Error
	if err != nil {
		log.Println("Error fetching triggers from DB:", err)
		return
	}

	fmt.Println("Triggers from database:", storedTriggers)

	// Now, storedTriggers will correctly contain the results from the database.
	for _, trigger := range storedTriggers {
		fmt.Println("Scheduling trigger:", trigger.ID, "with schedule:", trigger.Schedule)
		triggerCopy := trigger

		err := cron.AddFunc(trigger.Schedule, func() {
			executeScheduledTrigger(db, redisClient, triggerCopy)
		})
		if err != nil {
			log.Println("Failed to schedule trigger:", err)
		}
	}

	cron.Start()
}

// Execute scheduled trigger
func executeScheduledTrigger(db *gorm.DB, redisClient *redis.Client, trigger models.Trigger) {
	event := models.EventLog{
		TriggerID:   trigger.ID,
		TriggeredAt: time.Now(),
		Payload:     trigger.Payload,
		Type:        "scheduled",
		Status:      "active",
	}

	db.Create(&event)

	// Cache in Redis for fast access
	eventJSON, _ := json.Marshal(event)
	redisClient.Set("event:"+trigger.ID, eventJSON, 2*time.Hour)

	log.Printf("Trigger %s executed", trigger.ID)
}

// package schedulers

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"sync"
// 	"time"

// 	"github.com/robfig/cron/v3"
// 	"gorm.io/gorm"
// 	"github.com/go-redis/redis/v8"
// 	"event-trigger-platform/internal/models"
// )

// type Scheduler struct {
// 	Cron      *cron.Cron
// 	Jobs      map[string]cron.EntryID // Store triggerID -> Job ID
// 	Mutex     sync.Mutex              // Prevent race conditions
// 	DB        *gorm.DB
// 	Redis     *redis.Client
// }

// // StartScheduler initializes the scheduler and continuously checks for updates
// func StartScheduler(db *gorm.DB, redisClient *redis.Client) *Scheduler {
// 	s := &Scheduler{
// 		Cron:  cron.New(),
// 		Jobs:  make(map[string]cron.EntryID),
// 		DB:    db,
// 		Redis: redisClient,
// 	}

// 	// Load initial triggers
// 	s.syncTriggers()

// 	// Start a goroutine to refresh triggers every 1 minute
// 	go func() {
// 		for {
// 			time.Sleep(60 * time.Second)
// 			s.syncTriggers()
// 		}
// 	}()

// 	// Start the cron scheduler
// 	s.Cron.Start()
// 	return s
// }

// // Sync database triggers with running jobs
// func (s *Scheduler) syncTriggers() {
// 	s.Mutex.Lock()
// 	defer s.Mutex.Unlock()

// 	var triggers []models.Trigger
// 	s.DB.Where("type = ?", "scheduled").Find(&triggers)

// 	activeTriggerIDs := make(map[string]bool)

// 	// Add or update triggers
// 	for _, trigger := range triggers {
// 		activeTriggerIDs[trigger.ID] = true

// 		// If trigger is new, add it
// 		if _, exists := s.Jobs[trigger.ID]; !exists {
// 			jobID, err := s.Cron.AddFunc(trigger.Schedule, func() {
// 				s.executeScheduledTrigger(trigger)
// 			})
// 			if err != nil {
// 				log.Println("Failed to schedule trigger:", err)
// 				continue
// 			}
// 			s.Jobs[trigger.ID] = jobID
// 			log.Printf("Scheduled new trigger: %s (%s)", trigger.ID, trigger.Schedule)
// 		}
// 	}

// 	// Remove deleted triggers
// 	for triggerID, jobID := range s.Jobs {
// 		if !activeTriggerIDs[triggerID] {
// 			s.Cron.Remove(jobID)
// 			delete(s.Jobs, triggerID)
// 			log.Printf("Removed trigger: %s", triggerID)
// 		}
// 	}
// }

// // Execute the scheduled trigger
// func (s *Scheduler) executeScheduledTrigger(trigger models.Trigger) {
// 	event := models.EventLog{
// 		TriggerID:   trigger.ID,
// 		TriggeredAt: time.Now(),
// 		Payload:     trigger.Payload,
// 		Type:        "scheduled",
// 		Status:      "active",
// 	}

// 	s.DB.Create(&event)

// 	// Cache in Redis for fast lookup
// 	eventJSON, _ := json.Marshal(event)
// 	ctx := context.Background()
// 	s.Redis.Set(ctx, "event:"+trigger.ID, eventJSON, 2*time.Hour)

// 	log.Printf("Executed scheduled trigger: %s", trigger.ID)
// }

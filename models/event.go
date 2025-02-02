package models

import (
	"time"
)

type EventLog struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TriggerID   string    `gorm:"not null"`
	TriggeredAt time.Time `gorm:"autoCreateTime"`
	Payload     string    `gorm:"type:jsonb"`
	Type        string    `gorm:"not null"`         // "scheduled" or "api"
	Status      string    `gorm:"default:'active'"` // "active", "archived"
}

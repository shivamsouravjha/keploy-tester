package models

import (
	"encoding/json"
	"time"
)

type EventLog struct {
	ID          string          `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TriggerID   string          `gorm:"not null"`
	TriggeredAt time.Time       `gorm:"autoCreateTime"`
	Payload     json.RawMessage `gorm:"type:jsonb" swaggertype:"object"`
	Type        string          `gorm:"not null"`         // "scheduled" or "api"
	Status      string          `gorm:"default:'active'"` // "active", "archived"
}

package models

import (
	"encoding/json"
	"time"
)

type Trigger struct {
	ID        string          `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Type      string          `gorm:"not null"` // "scheduled" or "api"
	Schedule  string          `gorm:"default:null"`
	Endpoint  *string         `gorm:"default:null"`
	Payload   json.RawMessage `gorm:"type:jsonb"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
}

type EventLog struct {
	ID          string          `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TriggerID   string          `gorm:"not null"`
	TriggeredAt time.Time       `gorm:"autoCreateTime"`
	Payload     json.RawMessage `gorm:"type:jsonb"`
	Type        string          `gorm:"not null"`         // "scheduled" or "api"
	Status      string          `gorm:"default:'active'"` // "active", "archived"
}

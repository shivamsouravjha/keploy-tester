package models

import (
	"time"
)

type Trigger struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Type      string    `gorm:"not null"` // "scheduled" or "api"
	Schedule  string    `gorm:"default:null"`
	Endpoint  *string   `gorm:"default:null"`
	Payload   string    `gorm:"type:jsonb"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	OneTime   bool      `json:"one_time"`
}

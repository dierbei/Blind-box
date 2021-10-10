package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32          `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

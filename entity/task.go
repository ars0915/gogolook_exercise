package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Name      *string        `json:"name" gorm:"type:varchar(255);not null;default:''"`
	Status    *uint8         `json:"status" gorm:"not null;default:0"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

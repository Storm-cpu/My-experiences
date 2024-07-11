package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	CreateAt time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `json:"delete_at" gorm:"index"`
}

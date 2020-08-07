package models

import "time"

// BaseModel 应用信息
type BaseModel struct {
	ID        uint `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string         `json:"username" gorm:"type:varchar(32);not null;uniqueIndex"`
	Password  string         `json:"-" gorm:"type:varchar(256);not null"` // json:"-" 确保密码不会被序列化
	Email     string         `json:"email" gorm:"type:varchar(128);uniqueIndex"`
	Avatar    string         `json:"avatar" gorm:"type:varchar(256)"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // 支持软删除
}

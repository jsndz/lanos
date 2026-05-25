package auth

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	UserName  string `gorm:"size:50;uniqueIndex;not null"`
	Email     string `gorm:"size:100;uniqueIndex;not null"`
	Password  string `gorm:"size:255;not null"`
	CreatedAt time.Time
}

type Session struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	Token     string    `gorm:"size:255;uniqueIndex;not null"`
	UserAgent string    `gorm:"size:255"`
	IPAddress string    `gorm:"size:45"`
	ExpiresAt time.Time `gorm:"index;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

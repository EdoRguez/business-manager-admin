package models

import "time"

type User struct {
	ID           int64     `gorm:"primaryKey"`
	CompanyID    int64     `gorm:"not null"`
	RoleID       int64     `gorm:"not null"`
	Email        string    `gorm:"not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	ModifiedAt   time.Time `gorm:"not null"`
}

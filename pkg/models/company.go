package models

import (
	"database/sql"
	"time"
)

type Company struct {
	ID              int64  `gorm:"primaryKey"`
	Name            string `gorm:"not null"`
	NameFormatUrl   string `gorm:"not null"`
	ImageUrl        sql.NullString
	PlanID          int64     `gorm:"not null"`
	LastPaymentDate time.Time `gorm:"not null"`
	CreatedAt       time.Time `gorm:"not null"`
	ModifiedAt      time.Time `gorm:"not null"`
}

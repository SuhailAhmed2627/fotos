package models

import (
	"time"
)

type UserRegistration struct {
	ID              uint      `gorm:"primary_key;AUTO_INCREMENT;not null;unique"`
	Username        string    `gorm:"default:null;"`
	Name            string    `gorm:"not null"`
	Email           string    `gorm:"not null;unique"`
	Password        string    `gorm:"not null"`
	EmailToken      string    `gorm:"type:varchar(255);not null"`
	IsEmailVerified bool      `gorm:"default:false;not null"`
	ForgotPwdToken  string    `gorm:"type:varchar(255);not null"`
	UserId          uint      `gorm:"default:null;"`
	CreatedAt       time.Time `gorm:"not null"`
	UpdatedAt       time.Time `gorm:"not null"`
}

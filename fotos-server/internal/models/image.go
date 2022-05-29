package models

import (
	"time"
)

type Image struct {
	ID            uint   `gorm:"primary_key;AUTO_INCREMENT;not null;unique"`
	Url           string `gorm:"not null;unique"`
	Key           string `gorm:"not null;unique"`
	Name          string `gorm:"not null"`
	Width         int    `gorm:"not null"`
	Height        int    `gorm:"not null"`
	FaceProcessed bool   `gorm:"default:false"`

	UploaderID   uint      `gorm:"not null"`
	Uploader     User      `gorm:"foreignkey:UploaderID"`
	PresentUsers []*User   `gorm:"many2many:image_present_users;"`
	EventID      uint      `gorm:"not null"`
	Event        Event     `gorm:"foreignkey:EventID"`
	ClickedAt    time.Time `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

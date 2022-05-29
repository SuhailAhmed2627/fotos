package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT;not null;unique"`
	Username  string `gorm:"default:null;"`
	Name      string `gorm:"default:null;"`
	Email     string `gorm:"not null;unique"`
	FaceUrl   string `gorm:"default:NEW_USER;"`
	UserToken string `gorm:"type:varchar(255);"`

	Events          []*Event `gorm:"many2many:user_events;"`
	PresentInImages []*Image `gorm:"many2many:image_present_users;"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

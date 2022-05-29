package models

import (
	"time"
)

type Face struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT;not null;unique"`
	AWSFaceID string `gorm:"not null;unique"`

	OwnerId uint
	Owner   User `gorm:"foreignKey:OwnerId"`
	EventId uint
	Event   Event `gorm:"foreignKey:EventId"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

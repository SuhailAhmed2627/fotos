package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT;not null;unique"`
	Name        string `gorm:"default:null;"`
	Description string `gorm:"type:varchar(255);not null"`
	Url         string `gorm:"type:varchar(255);unique"`

	CreatorID uint
	Creator   User     `gorm:"foreignkey:CreatorID"`
	Users     []*User  `gorm:"many2many:user_events;"`
	Images    []*Image `gorm:"foreignkey:EventID"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	e.Url = uuid.New().String()
	return
}

package models

import "gorm.io/gorm"

type Event struct {
	Content string

	Emotion string

	Category string

	UserID uint

	gorm.Model
}

type User struct {
	ID uint `gorm:"primaryKey"`

	Name     string
	Password string

	PortraitName string

	Events []Event `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	gorm.Model
}

package models

import "gorm.io/gorm"

type Input struct {
	ID uint `gorm:"primaryKey"`

	Content string

	Summary string

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

	Inputs []Input `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	gorm.Model
}

package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name string `gorm:"not null"`
}

//`gorm:"unique;type:varchar(50);not null"`

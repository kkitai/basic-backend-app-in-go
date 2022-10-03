package model

import "gorm.io/gorm"

// Telephone represents ORM struct for "telephones"
type Telephone struct {
	OwnerId int    `gorm:"not null"`
	Number  string `gorm:"unique;not null"`
	ICCId   int
	gorm.Model
}

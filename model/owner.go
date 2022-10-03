package model

import "gorm.io/gorm"

type Owner struct {
	Name    string
	Address string
	gorm.Model
}

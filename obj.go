package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string `gorm:"unique_index, not null"`
	LoginID  string
	Password string
	Name     string
}

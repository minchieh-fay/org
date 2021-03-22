package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey" json:"uid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//UUID     string `gorm:"unique_index, not null"`
	LoginID  string `gorm:"unique_index, not null"`
	Password string
	Name     string
}

type Tag struct {
	ID        uint `gorm:"primarykey" json:"gid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//UUID     string `gorm:"unique_index, not null"`
	Key   string
	Value string
	Uid   uint
}

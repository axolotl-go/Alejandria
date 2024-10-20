package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Name          string `gorm:"not null" json:"name"`
	Email         string `gorm:"unique;not null" json:"email"`
	Password_Hash string `gorm:"not null" json:"password_hash"`
}

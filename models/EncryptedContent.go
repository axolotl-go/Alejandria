package models

import "gorm.io/gorm"

type EncryptedContent struct {
	gorm.Model

	User_id int    `gorm:"unique;not null" json:"user_id"`
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
}

package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	//Soft Delete Implementation
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Todo struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  bool   `json:"status"`
	//Soft Delete Implementation
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

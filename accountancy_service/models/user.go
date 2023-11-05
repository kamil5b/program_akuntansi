package models

import (
	"gorm.io/gorm"
)

// User in this service
type User struct {
	gorm.Model
	Name string `json:"name"`
	Role string `json:"role"`
}

// All functions in this service is in User, this struct act like a bridge between Authorization and this service
type Account struct {
	AuthID uint `gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	User   User `json:"user"`
}

// NEXT UP : REPOSITORY FOR ACCOUNT

package app

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `json:"password" binding:"required,min=6,max=100"`
	Photo    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"column:email" json:"email,omitempty"`
	Password string `gorm:"column:password" json:"password"`
	Books    []*Book
}

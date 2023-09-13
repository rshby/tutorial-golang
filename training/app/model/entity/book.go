package entity

import (
	"time"
)

type Book struct {
	ID          int       `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	Title       string    `gorm:"column:title" json:"title,omitempty"`
	Description string    `gorm:"column:description" json:"description,omitempty"`
	Price       int       `gorm:"column:price" json:"price,omitempty"`
	Rating      int       `gorm:"rating" json:"rating,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UserId      uint      `gorm:"column:user_id" json:"user_id"`
}

package entity

import "time"

type Content struct {
	ID            int64     `json:"id,omitempty"`
	Title         string    `json:"title,omitempty"`
	PictureUrl    string    `json:"picture_url,omitempty"`
	TextFill      string    `json:"text_fill,omitempty"`
	Like          int       `json:"like,omitempty"`
	Dislike       int       `json:"dislike,omitempty"`
	AverageRating float64   `json:"average_rating,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	AccountId     int64     `json:"account_id,omitempty"`
}

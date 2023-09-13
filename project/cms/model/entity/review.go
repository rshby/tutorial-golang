package entity

import "time"

type Review struct {
	ID        int64     `json:"id,omitempty"`
	Rating    int       `json:"rating,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	AccountId int64     `json:"account_id,omitempty"`
	ContentId int64     `json:"content_id,omitempty"`
}

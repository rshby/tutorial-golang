package entity

import "time"

type User struct {
	ID         int64     `json:"id,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	IdentityId string    `json:"identity_id,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	Address    string    `json:"address,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

type Student struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

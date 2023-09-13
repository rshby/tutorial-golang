package entity

type Dislike struct {
	ID        int64 `json:"id,omitempty"`
	AccountId int64 `json:"account_id,omitempty"`
	ContentId int64 `json:"content_id,omitempty"`
}

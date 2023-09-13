package dto

// create request Like
type LikeRequest struct {
	AccountId int64 `json:"account_id" validate:"required"`
	ContentId int64 `json:"content_id" validate:"required"`
}

type LikeResponse struct {
	Account *AccountDetail `json:"account,omitempty"`
	Content *ContentDetail `json:"content,omitempty"`
}

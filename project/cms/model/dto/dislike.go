package dto

// create object request dislike
type DislikeRequest struct {
	AccountId int64 `json:"account_id" validate:"required"`
	ContentId int64 `json:"content_id" validate:"required"`
}

type DislikeResponse struct {
	Account *AccountDetail `json:"account,omitempty"`
	Content *ContentDetail `json:"content,omitempty"`
}

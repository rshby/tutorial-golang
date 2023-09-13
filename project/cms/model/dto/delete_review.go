package dto

// create object to request delete reviews
type DeleteReviewRequest struct {
	ReviewId  int64 `json:"review_id" validate:"required"`
	AccountId int64 `json:"account_id" validate:"required"`
	ContentId int64 `json:"content_id" validate:"required"`
}

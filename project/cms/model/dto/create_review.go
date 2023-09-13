package dto

// create obejct request insert review
type CreateReviewRequest struct {
	Rating    int    `json:"rating" validate:"required"`
	Comment   string `json:"comment" validate:"required"`
	AccountId int64  `json:"account_id" validate:"required"`
	ContentId int64  `json:"content_id" validate:"required"`
}

// object response insert review
type CreateReviewResponse struct {
	Id      int64          `json:"id,omitempty"`
	Review  *ReviewDetail  `json:"review,omitempty"`
	Content *ContentDetail `json:"content,omitempty"`
	Account *AccountDetail `json:"account,omitempty"`
}

type ReviewDetail struct {
	Id        int64  `json:"id,omitempty"`
	Rating    int    `json:"rating,omitempty"`
	Comment   string `json:"comment,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	AccountId int64  `json:"account_id,omitempty"`
	ContentId int64  `json:"content_id,omitempty"`
}

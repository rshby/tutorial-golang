package dto

// create object request create-content
type CreateContentRequest struct {
	AccountId  int64  `json:"account_id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	PictureUrl string `json:"picture_url" validate:"required"`
	TextFill   string `json:"text_fill" validate:"required"`
}

type CreateContentResponse struct {
	Content   *ContentDetail    `json:"content,omitempty"`
	CreatedBy *ContentCreatedBy `json:"created_by,omitempty"`
}

type ContentDetail struct {
	ID            int64   `json:"id,omitempty"`
	Title         string  `json:"title,omitempty"`
	PictureUrl    string  `json:"picture_url,omitempty"`
	TextFill      string  `json:"text_fill,omitempty"`
	Like          int     `json:"like"`
	Dislike       int     `json:"dislike"`
	AverageRating float64 `json:"average_rating"`
	CreatedAt     string  `json:"created_at,omitempty"`
	AccountId     int64   `json:"account_id,omitempty"`
}

type ContentCreatedBy struct {
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	FullName  string `json:"full_name,omitempty"`
}

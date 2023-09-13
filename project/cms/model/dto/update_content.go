package dto

// create object to request update content
type UpdateContentRequest struct {
	Id         int64  `json:"id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	PictureUrl string `json:"picture_url" validate:"required"`
	TextFill   string `json:"text_fill" validate:"required"`
}

type UpdateContentResponse struct {
	Id            int64   `json:"id,omitempty"`
	Title         string  `json:"title,omitempty"`
	PictureUrl    string  `json:"picture_url,omitempty"`
	TextFill      string  `json:"text_fill,omitempty"`
	Like          int     `json:"like"`
	Dislike       int     `json:"dislike"`
	AverageRating float64 `json:"average_rating"`
	CreatedAt     string  `json:"created_at"`
}

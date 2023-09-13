package dto

type UpdateBookRequest struct {
	Id          int    `json:"id" binding:"required,number"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Rating      int    `json:"rating" binding:"required,number"`
}

package dto

type InsertBookRequest2 struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Rating      int    `json:"rating" binding:"required,number"`
	UserId      int    `json:"user_id" binding:"required,number"`
}

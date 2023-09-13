package dto

type InsertBookRequest struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price" binding:"required,number"`
	SubTitle string `json:"sub_title" binding:"required"`
}

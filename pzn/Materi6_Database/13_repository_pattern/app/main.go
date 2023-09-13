package main

import (
	"13_repository_pattern/controller"
	"13_repository_pattern/entity"
)

func main() {
	comment1 := entity.Comment{
		Email:   "reosahobby@gmail.com",
		Comment: "Ini Komen Kedua",
	}

	// insert ke database
	controller.InsertComment(&comment1)

	// get comment by Id
	controller.GetCommentById(5)

	// get all data comment
	controller.GetAllComment()
}

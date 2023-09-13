package controller

import (
	"13_repository_pattern/connection"
	"13_repository_pattern/entity"
	"13_repository_pattern/repository"
	"context"
	"fmt"
)

// function untuk insert
func InsertComment(inputData *entity.Comment) {
	commentRepo := repository.NewCommentRepository(connection.ConnectDatabase())
	ctx := context.Background()
	result, err := commentRepo.Insert(ctx, *inputData)
	if err != nil {
		panic("Gagal Insert Controller " + err.Error())
	}

	fmt.Println(result)
}

// function untuk Get All Comment Data
func GetAllComment() {
	commentRepo := repository.NewCommentRepository(connection.ConnectDatabase())
	ctx := context.Background()

	result, err := commentRepo.FindAll(ctx)
	if err != nil {
		panic("Gagal Get All Comment Controller = " + err.Error())
	}

	fmt.Println(result)
}

// function untuk Get Comment By Id
func GetCommentById(id int32) {
	commentRepo := repository.NewCommentRepository(connection.ConnectDatabase())
	ctx := context.Background()

	result, err := commentRepo.FindById(ctx, id)
	if err != nil {
		panic("Gagal Get Comment By Id Controller " + err.Error())
	}

	fmt.Println(result.Id, ", ", result.Email, ", ", result.Comment)
}

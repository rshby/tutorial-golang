package service

import (
	"training/app/model/dto"
	"training/app/model/entity"
)

type IBookService interface {
	InsertBook(request *dto.InsertBookRequest2) (*entity.Book, error)
	GetAllBooks() ([]*entity.Book, error)
	GetBookById(id int) (*entity.Book, error)
	DeleteBookById(id int) error
	UpdateBookById(request *dto.UpdateBookRequest) (*entity.Book, error)
	GetByUserId(userId int) ([]*entity.Book, error)
}

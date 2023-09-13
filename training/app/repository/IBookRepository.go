package repository

import "training/app/model/entity"

type IBookRepository interface {
	InsertBook(entity *entity.Book) (*entity.Book, error)
	GetAllBooks() ([]*entity.Book, error)
	GetBookById(id int) (*entity.Book, error)
	DeleteBookById(entity *entity.Book) error
	UpdateBookById(id int, oldData *entity.Book, newData *entity.Book) (*entity.Book, error)
	GetByUserId(userId int) ([]*entity.Book, error)
}

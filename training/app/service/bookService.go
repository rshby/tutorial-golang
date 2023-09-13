package service

import (
	"errors"
	"fmt"
	"time"
	"training/app/model/dto"
	entity "training/app/model/entity"
	"training/app/repository"
)

type BookService struct {
	BookRepository *repository.BookRepository
	UserRepository *repository.UserRepository
}

// create function provider
func NewBookService(bookRepo *repository.BookRepository, userRepo *repository.UserRepository) *BookService {
	return &BookService{
		BookRepository: bookRepo,
		UserRepository: userRepo,
	}
}

// method insert book
func (b *BookService) InsertBook(request *dto.InsertBookRequest2) (*entity.Book, error) {
	// cek id
	user, err := b.UserRepository.GetUserById(request.UserId)
	fmt.Println("user:", user)
	if err != nil {
		fmt.Println("masuk ke sini..")
		return nil, errors.New("user not found in database")
	}

	entity := &entity.Book{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Rating:      request.Rating,
		CreatedAt:   time.Now(),
		UserId:      uint(request.UserId),
	}

	resultInsert, err := b.BookRepository.InsertBook(entity)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// success insert
	return resultInsert, nil
}

// method get all books
func (b *BookService) GetAllBooks() ([]*entity.Book, error) {
	// call procedure in repository
	books, err := b.BookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

// method to get book by id
func (b *BookService) GetBookById(id int) (*entity.Book, error) {
	book, err := b.BookRepository.GetBookById(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// method to delete data book by id
func (b *BookService) DeleteBookById(id int) error {
	// get data by id
	book, err := b.BookRepository.GetBookById(id)
	if err != nil {
		return err
	}

	// call procedure delete in repo
	err = b.BookRepository.DeleteBookById(book)
	return err
}

// method update book by Id
func (b *BookService) UpdateBookById(request *dto.UpdateBookRequest) (*entity.Book, error) {
	// get old data
	oldBook, err := b.BookRepository.GetBookById(request.Id)
	if err != nil {
		return nil, errors.New("record not found")
	}

	// ceate new data from request
	newData := &entity.Book{
		ID:          oldBook.ID,
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Rating:      request.Rating,
		CreatedAt:   oldBook.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	// call procedure update in repository
	newBook, err := b.BookRepository.UpdateBookById(request.Id, oldBook, newData)
	if err != nil {
		return nil, err
	}

	return newBook, nil
}

// method get book by user_id
func (b *BookService) GetByUserId(userId int) ([]*entity.Book, error) {
	// cek data user in repository layer
	_, err := b.UserRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	// call procedure GetByUserId in repository layer
	books, err := b.BookRepository.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	// success get books by user_id
	return books, nil
}

package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"training/app/model/entity"
)

type BookRepository struct {
	Db *gorm.DB
}

// create function provider
func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		Db: db,
	}
}

// method to create insert book
func (b *BookRepository) InsertBook(entity *entity.Book) (*entity.Book, error) {
	err := b.Db.Create(&entity).Error
	if err != nil {
		return nil, err
	}

	// success insert
	return entity, nil
}

// method to get all books
func (b *BookRepository) GetAllBooks() ([]*entity.Book, error) {
	var books []*entity.Book
	result := b.Db.Find(&books)

	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	return books, nil
}

// method to get book by id
func (b *BookRepository) GetBookById(id int) (*entity.Book, error) {
	var book *entity.Book
	result := b.Db.Find(&book, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	return book, nil
}

// method to delete book by id
func (b *BookRepository) DeleteBookById(entity *entity.Book) error {
	result := b.Db.Delete(&entity)
	if result.Error != nil {
		return result.Error
	}

	// success delete data book by id
	return nil
}

// method to update book by id
func (b *BookRepository) UpdateBookById(id int, oldData *entity.Book, newData *entity.Book) (*entity.Book, error) {
	oldData.Title = newData.Title
	oldData.Description = newData.Description
	oldData.Price = newData.Price
	oldData.Rating = newData.Rating
	oldData.UpdatedAt = newData.UpdatedAt

	result := b.Db.Save(&oldData)
	if result.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("cant update: %v", result.Error.Error()))
	}

	return oldData, nil
}

// method get data book by user_id
func (b *BookRepository) GetByUserId(userId int) ([]*entity.Book, error) {
	var books []*entity.Book
	result := b.Db.Find(&books, "user_id=?", userId)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return books, nil
}

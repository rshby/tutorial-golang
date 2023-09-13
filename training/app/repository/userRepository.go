package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"training/app/model/entity"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// method get all users
func (u *UserRepository) GetAllUsers() ([]*entity.User, error) {
	return nil, nil
}

// method get user by id
func (u *UserRepository) GetUserById(id int) (*entity.User, error) {
	var user entity.User
	result := u.DB.Find(&user, "id=?", id)
	if result.RowsAffected == 0 {
		return nil, errors.New("record user not found")
	}

	return &user, nil
}

// method to get user by email
func (u *UserRepository) GetByEmail(email string) (*entity.User, error) {
	var user *entity.User

	result := u.DB.Model(&entity.User{}).Where("email=?", email).First(&user)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return user, nil
}

// method insert new data user
func (u *UserRepository) InsertUser(entity *entity.User) (*entity.User, error) {
	result := u.DB.Create(&entity)

	if result.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("cant create new user: %v", result.Error))
	}

	return entity, nil
}

// method update user
func (u *UserRepository) UpdateUser(newData *entity.User) (*entity.User, error) {
	return nil, nil
}

package repository

import (
	"context"
	"errors"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB      *gorm.DB
	General *GeneralRepository[entity.User]
}

func NewUserRepository(db *gorm.DB, generalRepo *GeneralRepository[entity.User]) *UserRepository {
	return &UserRepository{
		DB:      db,
		General: generalRepo,
	}
}

// method get user information by user_id
func (u *UserRepository) GetUserInformationById(ctx context.Context, id int) (entity.User, error) {
	var user entity.User

	result := u.DB.WithContext(ctx).Model(&entity.User{}).Preload("Address.SubDistrict.District.City.Province").Preload("Account").Preload("Education.University").Where("users.id=?", id).First(&user)
	if result.Error != nil {
		// return with error data not found
		return entity.User{}, result.Error
	}

	// success get data user by Id
	// return
	return user, nil
}

// method get all Users information
func (u *UserRepository) GetAllUsersInformation(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	result := u.DB.WithContext(ctx).Model(&entity.User{}).Preload("Address.SubDistrict.District.City.Province").Preload("Account").Preload("Education.University").Find(&users)
	if result.RowsAffected == 0 {
		// return with error data not found
		return []entity.User{}, errors.New("record users not found")
	}

	// success get all users information
	// return
	return users, nil
}

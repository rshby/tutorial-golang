package repository

import "training/app/model/entity"

type IUserRepository interface {
	GetAllUsers() ([]*entity.User, error)
	GetUserById(id int) (*entity.User, error)
	InsertUser(entity *entity.User) (*entity.User, error)
	UpdateUser(newData *entity.User) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}

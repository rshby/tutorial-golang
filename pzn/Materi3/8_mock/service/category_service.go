package service

import (
	"Materi3/8_mock/entity"
	"Materi3/8_mock/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found!")
	} else {
		return category, nil
	}
}

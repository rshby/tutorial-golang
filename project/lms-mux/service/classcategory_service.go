package service

import (
	"lms-mux/repository"

	"github.com/go-playground/validator/v10"
)

type ClassCategoryService struct {
	ClassCategoryRepository *repository.ClassCategoryRepository
	Validate                *validator.Validate
}

package repository

import (
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type ClassCategoryRepository struct {
	DB      *gorm.DB
	General *GeneralRepository[entity.ClassCategory]
}

func NewClassCategoryRepository(db *gorm.DB, general *GeneralRepository[entity.ClassCategory]) *ClassCategoryRepository {
	return &ClassCategoryRepository{
		DB:      db,
		General: general,
	}
}

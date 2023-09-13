package repository

import (
	"Materi3/8_mock/entity"
)

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

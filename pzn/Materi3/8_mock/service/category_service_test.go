package service

import (
	"Materi3/8_mock/entity"
	"Materi3/8_mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// buat variabelnya dulu
var categoryRepository = &repository.CategoryRepositoryMock{
	Mock: mock.Mock{},
}

var categoryService = CategoryService{
	Repository: categoryRepository,
}

func TestCategoryServic_Get(t *testing.T) {
	//program mocknya
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

// buat function Get Success
func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)

}

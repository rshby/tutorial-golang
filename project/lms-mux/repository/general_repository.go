package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// buat sebuah generic struct
type GeneralRepository[T any] struct {
	DB     *gorm.DB
	Entity *T
}

// buat sebuah generic provider
func NewGeneralRepository[T any](db *gorm.DB, entity *T) *GeneralRepository[T] {
	return &GeneralRepository[T]{
		DB:     db,
		Entity: entity,
	}
}

// method Insert
func (g *GeneralRepository[T]) Insert(ctx context.Context, entity *T) (T, error) {
	var model T
	result := g.DB.WithContext(ctx).Model(&model).Create(&entity)
	if result.Error != nil {
		// return with error internal server
		return model, result.Error
	}

	// success insert data to database
	// return
	return *entity, nil
}

// generic method Update
func (g *GeneralRepository[T]) Update(ctx context.Context, id int, entity *T) (T, error) {
	var nilData T
	result := g.DB.WithContext(ctx).Model(&nilData).Where("id=?", id).Updates(&entity)
	if result.Error != nil {
		// return with error internal server
		return nilData, result.Error
	}

	// success update data by Id
	// return
	return *entity, nil
}

// generic method Delete
func (g *GeneralRepository[T]) Delete(ctx context.Context, entity *T) (string, error) {
	var model T
	result := g.DB.WithContext(ctx).Model(&model).Delete(&entity)
	if result.Error != nil {
		// return with error internal server
		return "", result.Error
	}

	// success delete data by Id
	return "success delete data by Id", nil
}

// generic method GetAll
func (g *GeneralRepository[T]) GetAll(ctx context.Context) ([]T, error) {
	var entities []T
	var model T
	var nilDatas []T

	result := g.DB.WithContext(ctx).Model(&model).Find(&entities)
	if result.RowsAffected == 0 {
		// return with error data not found
		return nilDatas, errors.New("record not found")
	}

	// success get all data entities
	// return
	return entities, nil
}

// generic method GetById
func (g *GeneralRepository[T]) GetById(ctx context.Context, id int) (T, error) {
	var model T
	var entity T
	result := g.DB.WithContext(ctx).Model(&model).Where("id=?", id).First(&entity)
	if result.Error != nil {
		// return with error data  not found
		return model, result.Error
	}

	// success get data by Id
	// return
	return entity, nil
}

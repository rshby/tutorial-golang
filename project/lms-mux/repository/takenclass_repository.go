package repository

import (
	"context"
	"errors"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type TakenClassRepository struct {
	DB      *gorm.DB
	General *GeneralRepository[entity.TakenClass]
}

func NewTakenClassRepository(db *gorm.DB, general *GeneralRepository[entity.TakenClass]) *TakenClassRepository {
	return &TakenClassRepository{
		DB:      db,
		General: general,
	}
}

// method Get data takenclass by Id
func (t *TakenClassRepository) GetById(ctx context.Context, id int) (entity.TakenClass, error) {
	var takenclass entity.TakenClass

	result := t.DB.WithContext(ctx).Model(&entity.TakenClass{}).Preload("User").Preload("Class").Where("taken_classes.id=?", id).First(&takenclass)
	if result.Error != nil {
		// return with error data not found
		return entity.TakenClass{}, result.Error
	}

	// success get data takenclass by Id
	// return
	return takenclass, nil
}

// method get all data takenclasses by user_id
func (t *TakenClassRepository) GetTakenClassByuserId(ctx context.Context, userId int) ([]entity.TakenClass, error) {
	var takenClasses []entity.TakenClass

	result := t.DB.WithContext(ctx).Model(&entity.TakenClass{}).Preload("User").Preload("Class").Where("taken_classes.user_id=?", userId).Find(&takenClasses)
	if result.RowsAffected == 0 {
		// return with error data not found
		return []entity.TakenClass{}, errors.New("record takenclass not found")
	}

	// success get data takenclass by user_id
	// return
	return takenClasses, nil
}

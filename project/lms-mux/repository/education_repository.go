package repository

import (
	"context"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type EducationRepository struct {
	DB      *gorm.DB
	General *GeneralRepository[entity.Education]
}

func NewEducationRepository(db *gorm.DB, generalRepo *GeneralRepository[entity.Education]) *EducationRepository {
	return &EducationRepository{
		DB:      db,
		General: generalRepo,
	}
}

// method get infromation education and university
func (e *EducationRepository) GetAllEducationInformation(ctx context.Context) ([]entity.Education, error) {
	educations := []entity.Education{}

	result := e.DB.WithContext(ctx).Model(&entity.Education{}).Preload("University").Find(&educations)
	if result.RowsAffected == 0 {
		// return with error data not found
		return []entity.Education{}, result.Error
	}

	// success get all data information education and university
	// return
	return educations, nil
}

// method get information education and university by educations_id
func (e *EducationRepository) GetEducationInformationById(ctx context.Context, educationId int) (entity.Education, error) {
	education := entity.Education{}

	result := e.DB.WithContext(ctx).Preload("University").Where("educations.id=?", educationId).First(&education)
	if result.Error != nil {
		// return with error data not found
		return entity.Education{}, result.Error
	}

	// success get education information by education_id
	// return
	return education, nil
}

package repository

import (
	"context"
	"errors"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type UniversityRepository struct {
	DB *gorm.DB
}

func NewUniversityRepository(db *gorm.DB) *UniversityRepository {
	return &UniversityRepository{
		DB: db,
	}
}

// method get university by Id
func (u *UniversityRepository) GetById(ctx context.Context, id int) (entity.University, error) {
	university := entity.University{}
	result := u.DB.WithContext(ctx).Model(&entity.University{}).Where("universities.id=?", id).First(&university)
	if result.Error != nil {
		// return with error data not found
		return entity.University{}, result.Error
	}

	// success get data university By Id
	// return
	return university, nil
}

// method get full university information by Id
func (u *UniversityRepository) GetFullUnivById(ctx context.Context, id int) (entity.University, error) {
	var university entity.University
	result := u.DB.WithContext(ctx).Model(&entity.University{}).Preload("Address.SubDistrict.District.City.Province").Where("universities.id=?", id).First(&university)
	if result.Error != nil {
		// return with error data not found
		return entity.University{}, result.Error
	}

	// success get full information university by id
	// return
	return university, nil
}

// method get all univesities
func (u *UniversityRepository) GetAll(ctx context.Context) ([]entity.University, error) {
	var universities []entity.University
	result := u.DB.WithContext(ctx).Model(&entity.University{}).Preload("Address.SubDistrict.District.City.Province").Find(&universities)
	if result.RowsAffected == 0 {
		// return with error data not found
		return []entity.University{}, errors.New("record not found")
	}

	// success get all data universities
	// return
	return universities, nil
}

// method get all educations by university_id
func (u *UniversityRepository) GetAllEducationsByUniversityId(ctx context.Context, universityId int) (entity.University, error) {
	var educations entity.University
	result := u.DB.WithContext(ctx).Model(&entity.University{}).Preload("Address.SubDistrict.District.City.Province").Preload("Educations").First(&educations)
	if result.Error != nil {
		// return with error data not found
		return entity.University{}, result.Error
	}

	// success get all educations by university_id
	// return
	return educations, nil
}

package repository

import (
	"context"
	"errors"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type DistrictRepository struct {
	DB *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) *DistrictRepository {
	return &DistrictRepository{
		DB: db,
	}
}

// method insert data district
func (d *DistrictRepository) Insert(ctx context.Context, district *entity.District) (entity.District, error) {
	result := d.DB.WithContext(ctx).Model(&entity.District{}).Create(&district)
	if result.Error != nil {
		// return with error internal server
		return entity.District{}, result.Error
	}

	// success insert data district
	// return
	return *district, nil
}

// method update data district by id
func (d *DistrictRepository) Update(ctx context.Context, newDistrict *entity.District) (entity.District, error) {
	result := d.DB.WithContext(ctx).Model(&entity.District{}).Where("districts.id=?", newDistrict.Id).Updates(&newDistrict)
	if result.Error != nil {
		// return with error internal server
		return entity.District{}, result.Error
	}

	// success update data district by id
	return *newDistrict, nil
}

// method delete data district by id
func (d *DistrictRepository) Delete(ctx context.Context, district *entity.District) (string, error) {
	result := d.DB.WithContext(ctx).Model(&entity.District{}).Where("districts.id=?", district.Id).Delete(&district)
	if result.Error != nil {
		// return with error internal server
		return "", result.Error
	}

	// success delete data district by id
	// return
	return "success delete data district", nil
}

// method get all data districts
func (d *DistrictRepository) GetAll(ctx context.Context) ([]entity.District, error) {
	var districts []entity.District
	result := d.DB.WithContext(ctx).Model(&entity.District{}).Find(&districts)
	if result.RowsAffected == 0 {
		// return with data not found
		return []entity.District{}, errors.New("record not found")
	}

	// success get all data districts
	// return
	return districts, nil
}

// method get data district by id
func (d *DistrictRepository) GetById(ctx context.Context, id int) (entity.District, error) {
	var district entity.District
	result := d.DB.WithContext(ctx).Model(&entity.District{}).Preload("City").Where("districts.id=?", id).First(&district)
	if result.Error != nil {
		// return with error data not found
		return entity.District{}, errors.New("record not found")
	}

	// success get data district by Id
	return district, nil
}

// method get all data districts by city_id
func (d *DistrictRepository) GetAllDistrictsByCityId(ctx context.Context, cityId int) (entity.City, error) {
	city := entity.City{}
	result := d.DB.WithContext(ctx).Model(&entity.City{}).Preload("Province").Preload("Districts").Where("cities.id=?", cityId).First(&city)
	if result.Error != nil {
		// return with data not found
		return entity.City{}, result.Error
	}

	// success get all districts by city_id
	// return
	return city, nil
}

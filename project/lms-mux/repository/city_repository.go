package repository

import (
	"context"
	"errors"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type CityRepository struct {
	DB *gorm.DB
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{
		DB: db,
	}
}

// method insert data city
func (c *CityRepository) Insert(ctx context.Context, city *entity.City) (entity.City, error) {
	result := c.DB.WithContext(ctx).Model(&entity.City{}).Create(&city)
	if result.Error != nil {
		// return with error internal server
		return entity.City{}, result.Error
	}

	// success insert data city
	// return
	return *city, nil
}

// method Update city by id
func (c *CityRepository) Update(ctx context.Context, newCity *entity.City) (entity.City, error) {
	result := c.DB.WithContext(ctx).Model(&entity.City{}).Where("id=?", newCity.Id).Updates(&newCity)
	if result.Error != nil {
		// return with error internal server
		return entity.City{}, result.Error
	}

	// success update data city by id
	// return
	return *newCity, nil
}

// method Delete city by id
func (c *CityRepository) Delete(ctx context.Context, city *entity.City) (string, error) {
	result := c.DB.WithContext(ctx).Model(&entity.City{}).Where("id=?", city.Id).Delete(&city)
	if result.Error != nil {
		// return with error internal server
		return "", result.Error
	}

	// success delete data city by id
	// return
	return "success delete data city by id", nil
}

// method get all city
func (c *CityRepository) GetAll(ctx context.Context) ([]entity.City, error) {
	cities := []entity.City{}
	result := c.DB.WithContext(ctx).Model(&entity.City{}).Preload("Province").Find(&cities)
	if result.RowsAffected == 0 {
		// return with error data not found
		return []entity.City{}, errors.New("record not found")
	}

	// success get all data cities
	// return
	return cities, nil
}

// method get city by id
func (c *CityRepository) GetById(ctx context.Context, id int) (entity.City, error) {
	city := entity.City{}
	result := c.DB.WithContext(ctx).Model(&entity.City{}).Preload("Province").Preload("Districts").Where("cities.id=?", id).First(&city)
	if result.Error != nil {
		// return with error data not found
		return entity.City{}, result.Error
	}

	// success get data city by id
	// return
	return city, nil
}

// method get all cities by province_id
func (c *CityRepository) GetAllCitiesByProvinceId(ctx context.Context, provinceId int) (entity.Province, error) {
	cities := entity.Province{}
	result := c.DB.WithContext(ctx).Model(&entity.Province{}).Preload("Cities").Where("provinces.id=?", provinceId).First(&cities)
	if result.Error != nil {
		// return with error data not found
		return entity.Province{}, result.Error
	}

	// success get all data cities by province_id
	// return
	return cities, nil
}

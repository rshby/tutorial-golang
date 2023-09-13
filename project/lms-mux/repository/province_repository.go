package repository

import (
	"context"
	"errors"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type ProvinceRepository struct {
	DB *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) *ProvinceRepository {
	return &ProvinceRepository{
		DB: db,
	}
}

// method Insert
func (p *ProvinceRepository) Insert(ctx context.Context, province *entity.Province) (entity.Province, error) {
	result := p.DB.WithContext(ctx).Model(&entity.Province{}).Create(&province)
	if result.Error != nil {
		// return with error
		return entity.Province{}, result.Error
	}

	// success insert data province
	// return
	return *province, nil
}

// method Update province data
func (p *ProvinceRepository) Update(ctx context.Context, newProvince *entity.Province) (entity.Province, error) {
	result := p.DB.WithContext(ctx).Model(&entity.Province{}).Where("id=?", newProvince.Id).Updates(&newProvince)
	if result.Error != nil {
		// return with error internal server
		return entity.Province{}, result.Error
	}

	// success update data province
	// return
	return *newProvince, nil
}

// method Delete province
func (p *ProvinceRepository) Delete(ctx context.Context, province *entity.Province) (string, error) {
	result := p.DB.WithContext(ctx).Model(&entity.Province{}).Where("id=?", province.Id).Delete(&province)
	if result.Error != nil {
		// return with error internal server
		return "", result.Error
	}

	// success delete data province
	// return
	return "success delete data province", nil
}

// method Get All provinces
func (p *ProvinceRepository) GetAll(ctx context.Context) ([]entity.Province, error) {
	provinces := []entity.Province{}
	result := p.DB.WithContext(ctx).Find(&provinces)
	if result.RowsAffected == 0 {
		// return with error not found
		return []entity.Province{}, errors.New("record not found")
	}

	// success get all data provinces
	// return
	return provinces, nil
}

// method Get Province by id
func (p *ProvinceRepository) GetById(ctx context.Context, id int) (entity.Province, error) {
	province := entity.Province{}
	//result := p.DB.WithContext(ctx).Where("id=?", id).First(&province)
	result := p.DB.WithContext(ctx).Preload("Cities").Where("provinces.id=?", id).First(&province)
	if result.Error != nil {
		// return with error not found
		return entity.Province{}, result.Error
	}

	// success get province by id
	// return
	return province, nil
}

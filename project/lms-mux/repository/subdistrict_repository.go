package repository

import (
	"context"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type SubDistrictRepository struct {
	DB *gorm.DB
}

func NewSubDistrictRepository(db *gorm.DB) *SubDistrictRepository {
	return &SubDistrictRepository{
		DB: db,
	}
}

// method insert data subdistrict
func (s *SubDistrictRepository) Insert(ctx context.Context, subdistrict *entity.SubDistrict) (entity.SubDistrict, error) {
	result := s.DB.WithContext(ctx).Model(&entity.SubDistrict{}).Create(&subdistrict)
	if result.Error != nil {
		// return with error internal server error
		return entity.SubDistrict{}, result.Error
	}

	// success insert data subdistrict
	// returns
	return *subdistrict, nil
}

// method update data subdistrict by id
func (s *SubDistrictRepository) Update(ctx context.Context, newSubDistrict *entity.SubDistrict) (entity.SubDistrict, error) {
	result := s.DB.WithContext(ctx).Model(&entity.SubDistrict{}).Where("sub_districts.id=?", newSubDistrict.Id).Updates(&newSubDistrict)
	if result.Error != nil {
		// return with error internal server error
		return entity.SubDistrict{}, result.Error
	}

	// success update data subdistrict by id
	// return
	return *newSubDistrict, nil
}

// method delete data subdistrict by id
func (s *SubDistrictRepository) Delete(ctx context.Context, subDistrict *entity.SubDistrict) (string, error) {
	result := s.DB.WithContext(ctx).Model(&entity.SubDistrict{}).Where("sub_districts.id=?", subDistrict.Id).Delete(&subDistrict)
	if result.Error != nil {
		// return with error internal server error
		return "", result.Error
	}

	// success delete data subdistrict by id
	// return
	return "success delete data subdistrict by id", nil
}

// method get data subdistrict by Id
func (s *SubDistrictRepository) GetById(ctx context.Context, id int) (entity.SubDistrict, error) {
	subDistrict := entity.SubDistrict{}
	result := s.DB.WithContext(ctx).Model(&entity.SubDistrict{}).Preload("District").Where("sub_districts.id=?", id).First(&subDistrict)
	if result.Error != nil {
		// return with error data not found
		return entity.SubDistrict{}, result.Error
	}

	// success get data subdistrict by Id
	// return
	return subDistrict, nil
}

// method get all data subdistricts by district_id
func (s *SubDistrictRepository) GetAllSubDistrictsByDistrictId(ctx context.Context, districtId int) (entity.District, error) {
	subDistricts := entity.District{}
	result := s.DB.WithContext(ctx).Model(&entity.District{}).Preload("City").Preload("SubDistricts").Where("districts.id=?", districtId).First(&subDistricts)
	if result.Error != nil {
		// return with error data not found
		return entity.District{}, result.Error
	}

	// success get all subdistricts by district_id
	// return
	return subDistricts, nil
}

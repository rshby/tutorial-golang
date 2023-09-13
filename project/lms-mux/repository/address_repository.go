package repository

import (
	"context"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type AddressRepository struct {
	General *GeneralRepository[entity.Address]
	DB      *gorm.DB
}

func NewAddressRepository(db *gorm.DB, general *GeneralRepository[entity.Address]) *AddressRepository {
	return &AddressRepository{
		DB:      db,
		General: general,
	}
}

// method GetById
func (a *AddressRepository) GetById(ctx context.Context, id int) (entity.Address, error) {
	address := entity.Address{}
	result := a.DB.WithContext(ctx).Model(&entity.Address{}).Preload("SubDistrict").Preload("User").Where("addresses.id=?", id).First(&address)
	if result.Error != nil {
		// return with error data not found
		return entity.Address{}, result.Error
	}

	// success get data address by Id
	// return
	return address, nil
}

// method get full address by id
func (a *AddressRepository) GetFullAddressById(ctx context.Context, id int) (entity.Address, error) {
	address := entity.Address{}
	result := a.DB.WithContext(ctx).Model(&entity.Address{}).Preload("SubDistrict.District.City.Province").Where("addresses.id=?", id).First(&address)
	if result.Error != nil {
		// return with error data not found
		return entity.Address{}, result.Error
	}

	// success get data full address by Id
	// return
	return address, nil
}

package service

import (
	"context"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"

	"github.com/go-playground/validator/v10"
)

type AddressService struct {
	GeneralRepository *repository.GeneralRepository[entity.Address]
	AddressRepository *repository.AddressRepository
	Validate          *validator.Validate
}

func NewAddressService(
	gRepo *repository.GeneralRepository[entity.Address],
	adRepo *repository.AddressRepository,
	validate *validator.Validate) *AddressService {
	return &AddressService{
		GeneralRepository: gRepo,
		AddressRepository: adRepo,
		Validate:          validate,
	}
}

// method insert address to database
func (a *AddressService) Insert(ctx context.Context, request *web.RequestAddressInsert) (entity.Address, error) {
	// validasi request
	err := a.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.Address{}, err
	}

	address := entity.Address{
		Street:        request.Street,
		SubDistrictId: request.SubDistrictId,
	}

	// jalankan perintah Insert yang ada di GeneralRepository
	result, err := a.GeneralRepository.Insert(ctx, &address)
	if err != nil {
		// return with error internal server
		return entity.Address{}, err
	}

	// success insert data address to database
	// return
	return result, nil
}

// method update address by Id
func (a *AddressService) Update(ctx context.Context, request *web.RequestAddressUpdate) (entity.Address, error) {
	// validasi request
	err := a.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.Address{}, err
	}

	// cek apakah data yang akan diupdate ada di database
	oldAddress, err := a.AddressRepository.GetById(ctx, request.Id)
	if err != nil {
		// return with error data not found
		return entity.Address{}, err
	}

	address := entity.Address{
		Id:            oldAddress.Id,
		Street:        request.Street,
		SubDistrictId: request.SubDistrictId,
	}

	// jalankan perintah Update yang ada di GeneralReposiory
	result, err := a.GeneralRepository.Update(ctx, address.Id, &address)
	if err != nil {
		// return with error internal server
		return entity.Address{}, err
	}

	// success update data address by Id
	// return
	return result, nil
}

// method delete address by Id
func (a *AddressService) Delete(ctx context.Context, id int) (string, error) {
	// cek apakah data yang akan diupdate ada di database
	oldAddress, err := a.AddressRepository.GetById(ctx, id)
	if err != nil {
		// return with error data not found
		return "", err
	}

	// jalankan perintah Delete yang ada di GeneralRepository
	result, err := a.GeneralRepository.Delete(ctx, &oldAddress)
	if err != nil {
		// return with error internal server
		return "", err
	}

	// success delete data address
	// return
	return result, nil
}

// method get all addresses
func (a *AddressService) GetAll(ctx context.Context) ([]entity.Address, error) {
	// jalankan perintah GetAll yang ada di GeneralRepository
	addresses, err := a.GeneralRepository.GetAll(ctx)
	if err != nil {
		// return with error data not found
		return []entity.Address{}, err
	}

	// success get all data addresses
	// return
	return addresses, nil
}

// method get address by Id
func (a *AddressService) GetById(ctx context.Context, id int) (entity.Address, error) {
	// jalankan perintah GetById yang ada di AddressRepository
	address, err := a.AddressRepository.GetById(ctx, id)
	if err != nil {
		// return with error not found
		return entity.Address{}, err
	}

	// success get data address by id
	// return
	return address, nil
}

// method get full adress by Id
func (a *AddressService) GetFullAddressById(ctx context.Context, id int) (web.ResponseFullAddress, error) {
	// jalankan perintah GetFullAddress yang ada di AddressRepository
	result, err := a.AddressRepository.GetFullAddressById(ctx, id)
	if err != nil {
		// return with error data not found
		return web.ResponseFullAddress{}, err
	}
	address := web.ResponseFullAddress{
		Id:          result.Id,
		Street:      result.Street,
		SubDistrict: result.SubDistrict.Name,
		ZipCode:     result.SubDistrict.ZipCode,
		District:    result.SubDistrict.District.Name,
		City:        result.SubDistrict.District.City.Name,
		Province:    result.SubDistrict.District.City.Province.Name,
	}

	// success get full address by Id
	// return
	return address, nil
}

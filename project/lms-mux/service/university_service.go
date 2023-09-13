package service

import (
	"context"
	"errors"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"

	"github.com/go-playground/validator/v10"
)

type UniversityService struct {
	GeneralRepository    *repository.GeneralRepository[entity.University]
	UniversityRepository *repository.UniversityRepository
	AddressRepository    *repository.GeneralRepository[entity.Address]
	Validate             *validator.Validate
}

func NewUniversityService(
	gRepo *repository.GeneralRepository[entity.University],
	univRepo *repository.UniversityRepository,
	addressRepo *repository.GeneralRepository[entity.Address],
	validate *validator.Validate) *UniversityService {
	return &UniversityService{
		GeneralRepository:    gRepo,
		UniversityRepository: univRepo,
		AddressRepository:    addressRepo,
		Validate:             validate,
	}
}

// method insert
func (u *UniversityService) Insert(ctx context.Context, request *web.RequestUniversityInsert) (entity.University, error) {
	// validasi request
	err := u.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.University{}, err
	}

	address := entity.Address{
		Street:        request.Street,
		SubDistrictId: request.SubDistrictId,
	}

	// insert Address menggunakan GeneralRepository
	addressInsert, err := u.AddressRepository.Insert(ctx, &address)
	if err != nil {
		// return with internal server
		return entity.University{}, err
	}

	university := entity.University{
		Name:      request.Name,
		AddressId: addressInsert.Id,
	}

	// jalankan perintah Insert yang ada di GeneralRepository
	result, err := u.GeneralRepository.Insert(ctx, &university)
	if err != nil {
		// return with error internal server
		return entity.University{}, err
	}

	// success insert university
	// return
	return result, nil
}

// method update university by Id
func (u *UniversityService) Update(ctx context.Context, request *web.RequestuniversityUpdate) (entity.University, error) {
	// validasi request
	err := u.Validate.Struct(*request)
	if err != nil {
		// return with error bad request
		return entity.University{}, err
	}

	// cek id univ apakah ada di database
	oldUniversty, err := u.UniversityRepository.GetById(ctx, request.Id)
	if err != nil {
		// return with error data not found
		return entity.University{}, errors.New("record university not found")
	}

	// get data oldAddress
	oldAddress, err := u.AddressRepository.GetById(ctx, oldUniversty.AddressId)
	if err != nil {
		// return data not found
		return entity.University{}, errors.New("record address not found")
	}

	address := entity.Address{
		Id:            oldAddress.Id,
		Street:        request.Street,
		SubDistrictId: request.SubDistrictId,
	}

	// update address
	_, err = u.AddressRepository.Update(ctx, oldAddress.Id, &address)
	if err != nil {
		// return with error internal server error
		return entity.University{}, err
	}

	university := entity.University{
		Id:        request.Id,
		Name:      request.Name,
		AddressId: oldAddress.Id,
	}

	// update university
	newUniversity, err := u.GeneralRepository.Update(ctx, request.Id, &university)
	if err != nil {
		// return with error internal server error
		return entity.University{}, err
	}

	// success update university
	// return
	return newUniversity, nil
}

// method Delete university
func (u *UniversityService) Delete(ctx context.Context, id int) (string, error) {
	// cek data yang akan dihapus apakah ada di database
	oldUniversity, err := u.GeneralRepository.GetById(ctx, id)
	if err != nil {
		// return with error data not found
		return "", errors.New("record university not found")
	}

	// cek data address
	oldAddress, err := u.AddressRepository.GetById(ctx, oldUniversity.AddressId)
	if err != nil {
		// return with error data not found
		return "", errors.New("record address not found")
	}

	// hapus university
	_, err = u.GeneralRepository.Delete(ctx, &oldUniversity)
	if err != nil {
		// return with error internal server
		return "", err
	}

	// hapus address yang terkait
	_, err = u.AddressRepository.Delete(ctx, &oldAddress)
	if err != nil {
		// return with error internal server error
		return "", err
	}

	// success delete university
	// return
	return "success delete university by id", nil
}

// method GetAll university
func (u *UniversityService) GetAll(ctx context.Context) ([]web.ResponseUniversity, error) {
	// jalankan perintah get all yang ada di UniversityRepository
	universities, err := u.UniversityRepository.GetAll(ctx)
	if err != nil {
		// return with error data not found
		return []web.ResponseUniversity{}, err
	}

	// success get all data universities
	var result []web.ResponseUniversity
	for _, data := range universities {
		result = append(result, web.ResponseUniversity{
			Id:   data.Id,
			Name: data.Name,
			Address: web.ResponseFullAddress{
				Id:          data.Address.Id,
				Street:      data.Address.Street,
				SubDistrict: data.Address.SubDistrict.Name,
				ZipCode:     data.Address.SubDistrict.ZipCode,
				District:    data.Address.SubDistrict.District.Name,
				City:        data.Address.SubDistrict.District.City.Name,
				Province:    data.Address.SubDistrict.District.City.Province.Name,
			},
		})
	}
	// return
	return result, nil
}

// method get data university by Id
func (u *UniversityService) GetById(ctx context.Context, id int) (web.ResponseUniversity, error) {
	// jalankan perintah get full univ by Id
	university, err := u.UniversityRepository.GetFullUnivById(ctx, id)
	if err != nil {
		// return with error data not found
		return web.ResponseUniversity{}, err
	}

	// success get full information university by Id
	result := web.ResponseUniversity{
		Id:   university.Id,
		Name: university.Name,
		Address: web.ResponseFullAddress{
			Id:          university.Address.Id,
			Street:      university.Address.Street,
			SubDistrict: university.Address.SubDistrict.Name,
			ZipCode:     university.Address.SubDistrict.ZipCode,
			District:    university.Address.SubDistrict.District.Name,
			City:        university.Address.SubDistrict.District.City.Name,
			Province:    university.Address.SubDistrict.District.City.Province.Name,
		},
	}

	// return
	return result, nil
}

// method get all educations by University_Id
func (u *UniversityService) GetAllEducationsByUniversityId(ctx context.Context, universityId int) (web.ResponseUniversity, error) {
	// cek data university_Id apakah ada di database
	university, err := u.GeneralRepository.GetById(ctx, universityId)
	if err != nil {
		// return with error data university not found
		return web.ResponseUniversity{}, errors.New("record university not found")
	}

	// jalankan perintah GetAllEducationsByUniversityId yg ada di repository
	univ, err := u.UniversityRepository.GetAllEducationsByUniversityId(ctx, university.Id)
	if err != nil {
		// return with error data not found
		return web.ResponseUniversity{}, err
	}

	// success get all educations by university_Id
	educations := web.ResponseUniversity{
		Id:   univ.Id,
		Name: univ.Name,
		Address: web.ResponseFullAddress{
			Id:          univ.Address.Id,
			Street:      univ.Address.Street,
			SubDistrict: univ.Address.SubDistrict.Name,
			ZipCode:     univ.Address.SubDistrict.ZipCode,
			District:    univ.Address.SubDistrict.District.Name,
			City:        univ.Address.SubDistrict.District.City.Name,
			Province:    univ.Address.SubDistrict.District.City.Province.Name,
		},
		Educations: univ.Educations,
	}

	// return
	return educations, nil
}

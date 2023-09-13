package service

import (
	"context"
	"errors"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"

	"github.com/go-playground/validator/v10"
)

type SubDistrictService struct {
	GeneralRepository     *repository.GeneralRepository[entity.SubDistrict]
	SubDistrictRepository *repository.SubDistrictRepository
	DistrictRepository    *repository.DistrictRepository
	Validate              *validator.Validate
}

func NewSubDistrictService(
	gRepo *repository.GeneralRepository[entity.SubDistrict],
	sdRepo *repository.SubDistrictRepository,
	dRepo *repository.DistrictRepository,
	validate *validator.Validate) *SubDistrictService {
	return &SubDistrictService{
		GeneralRepository:     gRepo,
		SubDistrictRepository: sdRepo,
		DistrictRepository:    dRepo,
		Validate:              validate,
	}
}

// method insert from general repo
func (s *SubDistrictService) Insert(ctx context.Context, request *web.RequestSubDistrictInsert) (entity.SubDistrict, error) {
	// validasi request
	err := s.Validate.Struct(*request)
	if err != nil {
		// return with error bad request
		return entity.SubDistrict{}, err
	}

	subDistrict := entity.SubDistrict{
		Name:       request.Name,
		ZipCode:    request.ZipCode,
		DistrictId: request.DistrictId,
	}

	// jalankan function insert yang ada di general_repository
	result, err := s.GeneralRepository.Insert(ctx, &subDistrict)
	if err != nil {
		// return with error internal server
		return entity.SubDistrict{}, err
	}

	// success insert data from general repo
	// return
	return result, nil
}

// method update data subdistrict
func (s *SubDistrictService) Update(ctx context.Context, request *web.RequestSubDistrictUpdate) (entity.SubDistrict, error) {
	// validasi request
	err := s.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.SubDistrict{}, err
	}

	// cek data yang ada di database apakah ada
	oldSubDistrict, err := s.SubDistrictRepository.GetById(ctx, request.Id)
	if err != nil {
		// return with error data not found
		return entity.SubDistrict{}, errors.New("record not found")
	}

	newSubDistrict := entity.SubDistrict{
		Id:         oldSubDistrict.Id,
		Name:       request.Name,
		ZipCode:    request.ZipCode,
		DistrictId: request.DistrictId,
	}

	// jalankan perintah Update yang ada di SubDistrictRepository
	result, err := s.SubDistrictRepository.Update(ctx, &newSubDistrict)
	if err != nil {
		// return with error internal server
		return entity.SubDistrict{}, err
	}

	// success update data subdistrict by Id
	// return
	return result, nil
}

// method delete data subdistrict
func (s *SubDistrictService) Delete(ctx context.Context, id int) (string, error) {
	// cek apakah data yang akan dihapus ada di database
	oldSubDistrict, err := s.SubDistrictRepository.GetById(ctx, id)
	if err != nil {
		// return with error not found
		return "", err
	}

	// jalankan perintah delete yang ada di SubDistrictRepository
	result, err := s.SubDistrictRepository.Delete(ctx, &oldSubDistrict)
	if err != nil {
		// return with error internal server error
		return "", err
	}

	// success delete data subdistrict by Id
	// return
	return result, nil
}

// method get all data subdistricts
func (s *SubDistrictService) GetAll(ctx context.Context) ([]entity.SubDistrict, error) {
	// jalankan perintah getall dari general repository
	subdistricts, err := s.GeneralRepository.GetAll(ctx)
	if err != nil {
		// return with error data not found
		return []entity.SubDistrict{}, err
	}

	// success get all subdistricts
	// return
	return subdistricts, nil
}

// method get data subdistrict by Id
func (s *SubDistrictService) GetById(ctx context.Context, id int) (entity.SubDistrict, error) {
	// jalankan perintah getbyid yang ada di SubDistrictRepository
	subdistrict, err := s.SubDistrictRepository.GetById(ctx, id)
	if err != nil {
		// return with error data not found
		return entity.SubDistrict{}, err
	}

	// success get data subdistrict by id
	// return
	return subdistrict, nil
}

// method get all data subdistricts by district_id
func (s *SubDistrictService) GetByDistrictId(ctx context.Context, districtId int) (entity.District, error) {
	// cek dulu district_id apakah ada di database
	district, err := s.DistrictRepository.GetById(ctx, districtId)
	if err != nil {
		// return with error data not found
		return entity.District{}, errors.New("data district not found")
	}

	// jalankan perintah GetByDistrictId yang ada di SubDistrictRepository
	subDistricts, err := s.SubDistrictRepository.GetAllSubDistrictsByDistrictId(ctx, district.Id)
	if err != nil {
		// return with error data not found
		return entity.District{}, err
	}

	// success get all data subdistricts by district_id
	// return
	return subDistricts, nil
}

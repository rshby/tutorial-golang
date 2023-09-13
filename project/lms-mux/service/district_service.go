package service

import (
	"context"
	"errors"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"

	"github.com/go-playground/validator/v10"
)

type DistrictService struct {
	DistrictRepository *repository.DistrictRepository
	CityRepository     *repository.CityRepository
	Validate           *validator.Validate
}

func NewDistrictService(districtRepo *repository.DistrictRepository, cityRepo *repository.CityRepository, validate *validator.Validate) *DistrictService {
	return &DistrictService{
		DistrictRepository: districtRepo,
		CityRepository:     cityRepo,
		Validate:           validate,
	}
}

// method insert data dictrict
func (d *DistrictService) Insert(ctx context.Context, request *web.RequestDistrictInsert) (entity.District, error) {
	// validasi request
	err := d.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.District{}, err
	}

	district := entity.District{
		Name:   request.Name,
		CityId: request.CityId,
	}

	// jalankan perintah insert yang ada di repository
	result, err := d.DistrictRepository.Insert(ctx, &district)
	if err != nil {
		// return with error internal server error
		return entity.District{}, err
	}

	// success insert data district
	// return
	return result, nil
}

// method update data district by Id
func (d *DistrictService) Update(ctx context.Context, request *web.RequestDistrictUpdate) (entity.District, error) {
	// validasi request
	err := d.Validate.Struct(*request)
	if err != nil {
		return entity.District{}, err
	}

	// cek dulu apakah data yang akan diupdate ada di database
	oldDistrict, err := d.DistrictRepository.GetById(ctx, request.Id)
	if err != nil {
		// return with data not found
		return entity.District{}, err
	}

	newDistrict := entity.District{
		Id:     oldDistrict.Id,
		Name:   request.Name,
		CityId: request.CityId,
	}

	// jalankan perintah update yang ada di repository
	result, err := d.DistrictRepository.Update(ctx, &newDistrict)
	if err != nil {
		// return with error internal server
		return entity.District{}, err
	}

	// success update data district by Id
	// return
	return result, nil
}

// method delete data district by Id
func (d *DistrictService) Delete(ctx context.Context, id int) (string, error) {
	// cek data yang akan didelete apakah ada di database
	oldDistrict, err := d.DistrictRepository.GetById(ctx, id)
	if err != nil {
		// return with error data not found
		return "", err
	}

	// jalankan perintah delete yang ada di repository
	result, err := d.DistrictRepository.Delete(ctx, &oldDistrict)
	if err != nil {
		// return with error internal server
		return "", err
	}

	// success delete data district by Id
	// return
	return result, nil
}

// method get all data districts
func (d *DistrictService) GetAll(ctx context.Context) ([]entity.District, error) {
	// jalankan perintah get all yang ada di DistrictRepository
	districts, err := d.DistrictRepository.GetAll(ctx)
	if err != nil {
		// return with error not found
		return []entity.District{}, err
	}

	// success get all data districts
	// return
	return districts, nil
}

// method get data district by Id
func (d *DistrictService) GetById(ctx context.Context, id int) (entity.District, error) {
	district, err := d.DistrictRepository.GetById(ctx, id)
	if err != nil {
		// return with error not found
		return entity.District{}, err
	}

	// success get data district by Id
	return district, nil
}

// method get all data districts by city_id
func (d *DistrictService) GetAllDistrictsByCityId(ctx context.Context, cityId int) (entity.City, error) {
	// cek dulu city_id apakah ada di database
	city, err := d.CityRepository.GetById(ctx, cityId)
	if err != nil {
		// return with error data not found
		return entity.City{}, errors.New("data city not found")
	}

	// jalankan perintah get all districts by city_id yang ada di DistrictRepository
	districts, err := d.DistrictRepository.GetAllDistrictsByCityId(ctx, city.Id)
	if err != nil {
		// return with data not found
		return entity.City{}, err
	}

	// success get all data districts by city_id
	// return
	return districts, nil
}

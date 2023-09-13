package service

import (
	"context"
	"errors"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"

	"github.com/go-playground/validator/v10"
)

type CityService struct {
	CityRepository     *repository.CityRepository
	ProvinceRepository *repository.ProvinceRepository
	Validate           *validator.Validate
}

func NewCityService(cityRepo *repository.CityRepository, provinceRepo *repository.ProvinceRepository, validate *validator.Validate) *CityService {
	return &CityService{
		CityRepository:     cityRepo,
		ProvinceRepository: provinceRepo,
		Validate:           validate,
	}
}

// method Insert data city
func (c *CityService) Insert(ctx context.Context, request *web.RequestCityInsert) (entity.City, error) {
	// validasi request
	err := c.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.City{}, err
	}

	city := entity.City{
		Name:       request.Name,
		ProvinceId: request.ProvinceId,
	}

	// jalankan perintah insert yang ada di repository
	result, err := c.CityRepository.Insert(ctx, &city)
	if err != nil {
		// return with error internal server
		return entity.City{}, err
	}

	// success insert data city
	// return
	return result, nil
}

// method update data city by id
func (c *CityService) Update(ctx context.Context, request *web.RequestCityUpdate) (entity.City, error) {
	// validasi request body
	err := c.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.City{}, err
	}

	// cek data ke database -> apakah data yang akan diupdate ada
	oldCity, err := c.CityRepository.GetById(ctx, request.Id)
	if err != nil {
		// return with error data not found
		return entity.City{}, errors.New("record not found")
	}

	newCity := entity.City{
		Id:         oldCity.Id,
		Name:       request.Name,
		ProvinceId: request.ProvinceId,
	}

	// jalankan perintah update city yang ada di repository
	result, err := c.CityRepository.Update(ctx, &newCity)
	if err != nil {
		// return with error internal server
		return entity.City{}, err
	}

	// success update data city by id
	// return
	return result, nil
}

// method delete city by id
func (c *CityService) Delete(ctx context.Context, id int) (string, error) {
	// cek data di database -> data yang akan dihapus ada atau tidak
	oldCity, err := c.CityRepository.GetById(ctx, id)
	if err != nil {
		// return with error data not found
		return "", errors.New("record not found")
	}

	// jalankan perintah delete city yang ada di repository
	result, err := c.CityRepository.Delete(ctx, &oldCity)
	if err != nil {
		// return with error internal server error
		return "", err
	}

	// success delete data city by id
	// return
	return result, nil
}

// method get all cites
func (c *CityService) GetAll(ctx context.Context) ([]entity.City, error) {
	// jalankan perintah get all yang ada di repository
	cities, err := c.CityRepository.GetAll(ctx)
	if err != nil {
		// return with error not found
		return []entity.City{}, err
	}

	// success get all data cities
	// return
	return cities, nil
}

// method get city by id
func (c *CityService) GetById(ctx context.Context, id int) (entity.City, error) {
	// jalankan perintah get city by Id yang ada di repository
	city, err := c.CityRepository.GetById(ctx, id)
	if err != nil {
		// return with error data not found
		return entity.City{}, err
	}

	// success get data city by id
	// return
	return city, nil
}

// method get all cities by province id
func (c *CityService) GetAllCitiesByProvinceId(ctx context.Context, provinceId int) (entity.Province, error) {
	// cek data province dengan province_id apakah ada di database
	_, err := c.ProvinceRepository.GetById(ctx, provinceId)
	if err != nil {
		// return with error province not found
		return entity.Province{}, errors.New("province not found")
	}

	// jalankan perintah get all cities by province id yang ada di repository
	cities, err := c.CityRepository.GetAllCitiesByProvinceId(ctx, provinceId)
	if err != nil {
		// return with error data not found
		return entity.Province{}, err
	}

	// success get all cities by province_id
	// return
	return cities, nil
}

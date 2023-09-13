package service

import (
	"context"
	"errors"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"

	"github.com/go-playground/validator/v10"
)

type ProvinceService struct {
	ProvinceRepository *repository.ProvinceRepository
	Validate           *validator.Validate
}

func NewProvinceService(provinceRepo *repository.ProvinceRepository, validate *validator.Validate) *ProvinceService {
	return &ProvinceService{
		ProvinceRepository: provinceRepo,
		Validate:           validate,
	}
}

// method Insert province
func (p *ProvinceService) Insert(ctx context.Context, request *web.RequestProvinceInsert) (entity.Province, error) {
	// validasi request
	err := p.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.Province{}, err
	}

	province := entity.Province{
		Name: request.Name,
	}

	// jalankan perintah Insert province yang ada di repository
	response, err := p.ProvinceRepository.Insert(ctx, &province)
	if err != nil {
		// return with error internal server
		return entity.Province{}, err
	}

	// success insert province
	// return
	return response, nil
}

// method Update province
func (p *ProvinceService) Update(ctx context.Context, request *web.RequestProvinceUpdate) (entity.Province, error) {
	// validasi request
	err := p.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.Province{}, err
	}

	// cek dulu apakah data yang akan diupdate ada di database
	oldProvince, err := p.ProvinceRepository.GetById(ctx, request.Id)
	if err != nil {
		// return with error not found
		return entity.Province{}, errors.New("record not found")
	}

	newProvince := entity.Province{
		Id:   oldProvince.Id,
		Name: request.Name,
	}

	// jalankan perintah update province yang ada di repository
	response, err := p.ProvinceRepository.Update(ctx, &newProvince)
	if err != nil {
		// return with error internal server
		return entity.Province{}, err
	}

	// success update data province
	// return
	return response, nil
}

// method Delete province
func (p *ProvinceService) Delete(ctx context.Context, id int) (string, error) {
	// cek data yang akan dihapus apakah ada di database
	province, err := p.ProvinceRepository.GetById(ctx, id)
	if err != nil {
		// return with error data  not found
		return "", errors.New("record not found")
	}

	// jalankan perintah Delete province yang ada di repository
	response, err := p.ProvinceRepository.Delete(ctx, &province)
	if err != nil {
		// return with error internal server
		return "", err
	}

	// success delete data province by id
	// return
	return response, nil
}

// method Get All provinces
func (p *ProvinceService) GetAll(ctx context.Context) ([]entity.Province, error) {
	// jalankan perintah GetAll provinces yang ada di repository
	provinces, err := p.ProvinceRepository.GetAll(ctx)
	if err != nil {
		// return with error data not found
		return []entity.Province{}, err
	}

	// success get all data provinces
	// return
	return provinces, nil
}

// method Get province By Id
func (p *ProvinceService) GetById(ctx context.Context, id int) (entity.Province, error) {
	// jalankan perintah Get province by Id yang ada di repository
	province, err := p.ProvinceRepository.GetById(ctx, id)
	if err != nil {
		// return with error not found
		return entity.Province{}, err
	}

	// success get data province by id
	// return
	return province, nil
}

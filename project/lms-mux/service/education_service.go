package service

import (
	"context"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"

	"github.com/go-playground/validator/v10"
)

type EducationService struct {
	EducationRepository *repository.EducationRepository
	Validate            *validator.Validate
}

func NewEducationService(educationRepo *repository.EducationRepository, validate *validator.Validate) *EducationService {
	return &EducationService{
		EducationRepository: educationRepo,
		Validate:            validate,
	}
}

// method insert
func (e *EducationService) Insert(ctx context.Context, request *web.RequestEducationInsert) (entity.Education, error) {
	// validasi request
	err := e.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.Education{}, err
	}

	// tambahkan object input data
	education := entity.Education{
		Major:        request.Major,
		Level:        request.Level,
		UniversityId: request.UniversityId,
	}

	// jalankan perintah Insert yang ada di repository
	result, err := e.EducationRepository.General.Insert(ctx, &education)
	if err != nil {
		// return with error internal server error
		return entity.Education{}, err
	}

	// success insert data education
	// return
	return result, nil
}

// method update education
func (e *EducationService) Update(ctx context.Context, request *web.RequestEducationUpdate) (entity.Education, error) {
	// cek data id apakah ada di database
	oldEducation, err := e.EducationRepository.General.GetById(ctx, request.Id)
	if err != nil {
		// return with error data not found
		return entity.Education{}, err
	}

	// validasi request
	err = e.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.Education{}, err
	}

	// buat newEducation
	newEducation := entity.Education{
		Id:           oldEducation.Id,
		Major:        request.Major,
		Level:        request.Level,
		UniversityId: request.UniversityId,
	}

	// jalankan perintah Update yang ada di repository
	result, err := e.EducationRepository.General.Update(ctx, oldEducation.Id, &newEducation)
	if err != nil {
		// return with error internal server
		return entity.Education{}, err
	}

	// success update data education
	// return
	return result, nil
}

// method delete data education by education_id
func (e *EducationService) Delete(ctx context.Context, id int) (string, error) {
	// cek data apakah id yang akan dihapus ada di database
	education, err := e.EducationRepository.General.GetById(ctx, id)
	if err != nil {
		// return with error data not found
		return "", err
	}

	// jalankan perintah delete yang ada di repository
	result, err := e.EducationRepository.General.Delete(ctx, &education)
	if err != nil {
		// return with error internal server
		return "", err
	}

	// success delete data
	// return
	return result, nil
}

// method get all education with university
func (e *EducationService) GetAll(ctx context.Context) ([]entity.Education, error) {
	// jalankan perintah get all yang ada di repository
	educations, err := e.EducationRepository.GetAllEducationInformation(ctx)
	if err != nil {
		// return with error not found
		return []entity.Education{}, err
	}

	// success get all educations with university
	// return
	return educations, nil
}

// method get data education with university by education_Id
func (e *EducationService) GetById(ctx context.Context, id int) (entity.Education, error) {
	// jalankan perintah get by id yang ada di repository
	education, err := e.EducationRepository.GetEducationInformationById(ctx, id)
	if err != nil {
		// return with error not found
		return entity.Education{}, err
	}

	// success get data education with university by education_Id
	// return
	return education, nil
}

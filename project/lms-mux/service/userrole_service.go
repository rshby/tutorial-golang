package service

import (
	"context"
	"errors"
	"lms-mux/helper"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserRoleService struct {
	UserRoleRepository *repository.UserRoleRepository
	Validate           *validator.Validate
}

func NewUserRoleService(userRoleRepo *repository.UserRoleRepository, validate *validator.Validate) *UserRoleService {
	return &UserRoleService{
		UserRoleRepository: userRoleRepo,
		Validate:           validate,
	}
}

// method insert userRole
func (u *UserRoleService) Insert(ctx context.Context, request *web.RequestUserRoleInsert) (entity.UserRole, error) {
	// validasi
	err := u.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.UserRole{}, err
	}

	// cek dulu apakah user_id yang ada sudah memiliki role yang diinput
	_, err = u.UserRoleRepository.GetByUserIdAndRoleId(ctx, request.UserId, request.RoleId)
	if err == nil {
		// return with error user already have role
		return entity.UserRole{}, errors.New("user already have this role")
	}

	// buat input
	userRole := entity.UserRole{
		UserId: request.UserId,
		RoleId: request.RoleId,
	}

	// jalankan perintah Insert yang ada di repository
	result, err := u.UserRoleRepository.General.Insert(ctx, &userRole)
	if err != nil {
		// return with error internal server
		return entity.UserRole{}, err
	}

	// get by Id
	response, err := u.UserRoleRepository.GetById(ctx, result.Id)
	if err != nil {
		// return with error not found
		return entity.UserRole{}, errors.New("record not found")
	}

	response.User.Gender = helper.ParseGender(response.User.Gender)

	// success insert
	return response, nil
}

// method update user-role by Id
func (u *UserRoleService) Update(ctx context.Context, request *web.RequestUserRoleUpdate) (entity.UserRole, error) {
	// validasi request
	err := u.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return entity.UserRole{}, err
	}

	// cek ke database apakah data user-role yang akan diupdate ada
	_, err = u.UserRoleRepository.General.GetById(ctx, request.Id)
	if err != nil {
		// return with data not found
		return entity.UserRole{}, errors.New("record user-role not found")
	}

	// buat input
	userRole := entity.UserRole{
		Id:     request.Id,
		UserId: request.UserId,
		RoleId: request.RoleId,
	}

	// jalankan perintah update yang ada di repository general
	result, err := u.UserRoleRepository.General.Update(ctx, userRole.Id, &userRole)
	if err != nil {
		// return with error internal server
		return entity.UserRole{}, err
	}

	// success update data user-role by Id
	responseData, err := u.UserRoleRepository.GetById(ctx, result.Id)
	if err != nil {
		// return with error data not found
		return entity.UserRole{}, errors.New("record user-role not found")
	}

	// success get new data
	// return
	return responseData, nil
}

// method delete user-role by Id
func (u *UserRoleService) Delete(ctx context.Context, id int) (string, error) {
	// cek data yang akan didelete apakah ada di database
	logrus.WithField("event", "UserRoleService-Delete").Info("proses get data user-role by Id")
	oldUserRole, err := u.UserRoleRepository.General.GetById(ctx, id)
	if err != nil {
		// return with error
		logrus.WithField("event", "UserRoleService-Delete").WithField("error", err.Error()).Info("ditemukan error")
		return "", errors.New("record user-role not found")
	}

	// jalankan perintah delete yang ada di repository
	result, err := u.UserRoleRepository.General.Delete(ctx, &oldUserRole)
	if err != nil {
		// return with internal server
		return "", err
	}

	// success delete data user-role by Id
	// return
	return result, nil
}

// method get all data user-roles
func (u *UserRoleService) GetAll(ctx context.Context) ([]entity.UserRole, error) {
	// jalankan perintah getall yang ada di repository
	userRoles, err := u.UserRoleRepository.GetAll(ctx)
	if err != nil {
		// jika data tidak ditemukan -> return with error
		return []entity.UserRole{}, err
	}

	// success get all data user-roles
	var response []entity.UserRole
	wg := &sync.WaitGroup{}
	mtx := &sync.RWMutex{}

	for _, userRole := range userRoles {
		wg.Add(1)
		go func(data entity.UserRole) {
			mtx.Lock()
			response = append(response, entity.UserRole{
				Id: data.Id,
				User: &entity.User{
					Id:        data.User.Id,
					FirstName: data.User.FirstName,
					LastName:  data.User.LastName,
					Gender:    helper.ParseGender(data.User.Gender),
					BirthDate: data.User.BirthDate,
				},
				Role: &entity.Role{
					Id:   data.Role.Id,
					Name: data.Role.Name,
				},
			})
			mtx.Unlock()
			wg.Done()
		}(userRole)
	}

	// return
	wg.Wait()
	return response, nil
}

// metod get data user-role by id
func (u *UserRoleService) GetById(ctx context.Context, id int) (entity.UserRole, error) {
	// jalankan perintah get by id yang ada di repository
	userRole, err := u.UserRoleRepository.GetById(ctx, id)
	if err != nil {
		// return wit error data not found
		return entity.UserRole{}, errors.New("record user-role not found")
	}

	response := entity.UserRole{
		Id: userRole.Id,
		User: &entity.User{
			Id:        userRole.User.Id,
			FirstName: userRole.User.FirstName,
			LastName:  userRole.User.LastName,
			Gender:    helper.ParseGender(userRole.User.Gender),
			BirthDate: userRole.User.BirthDate,
		},
		Role: &entity.Role{
			Id:   userRole.Role.Id,
			Name: userRole.Role.Name,
		},
	}

	// return
	return response, nil
}

// method get data user-role by user_id
func (u *UserRoleService) GetByUserId(ctx context.Context, userId int) ([]entity.UserRole, error) {
	// jalankan perintah getbyuserId yang ada di repository
	userRoles, err := u.UserRoleRepository.GetUserRolesByUserId(ctx, userId)
	if err != nil {
		// return with error data not found
		return []entity.UserRole{}, err
	}

	// success get data user-role by userId
	var response []entity.UserRole
	for _, data := range userRoles {
		response = append(response, entity.UserRole{
			Id: data.Id,
			User: &entity.User{
				Id:        data.User.Id,
				FirstName: data.User.FirstName,
				LastName:  data.User.LastName,
				Gender:    helper.ParseGender(data.User.Gender),
				BirthDate: data.User.BirthDate,
			},
			Role: &entity.Role{
				Id:   data.Role.Id,
				Name: data.Role.Name,
			},
		})
	}

	// return
	return response, nil
}

// method get data user-role by role_id
func (u *UserRoleService) GetByRoleId(ctx context.Context, roleId int) ([]entity.UserRole, error) {
	// jalankan perintah getby role_id
	userRoles, err := u.UserRoleRepository.GetByRoleId(ctx, roleId)
	if err != nil {
		// return with error data not found
		return []entity.UserRole{}, err
	}

	// success get data user-role by role_id
	var response []entity.UserRole
	for _, data := range userRoles {
		response = append(response, entity.UserRole{
			Id: data.Id,
			User: &entity.User{
				Id:        data.User.Id,
				FirstName: data.User.FirstName,
				LastName:  data.User.LastName,
				Gender:    helper.ParseGender(data.User.Gender),
				BirthDate: data.User.BirthDate,
			},
			Role: &entity.Role{
				Id:   data.Role.Id,
				Name: data.Role.Name,
			},
		})
	}

	// return
	return response, nil
}

// method get data user-role by user_id and role-id
func (u *UserRoleService) GetByUserIdAndRoleId(ctx context.Context, userId int, roleId int) (entity.UserRole, error) {
	// jalankan perintah get by user_id and role_id yang ada di repository
	userRole, err := u.UserRoleRepository.GetByUserIdAndRoleId(ctx, userId, roleId)
	if err != nil {
		// return with error data not found
		return entity.UserRole{}, err
	}

	// success get data user-role by user_id and role_id
	response := entity.UserRole{
		Id: userRole.Id,
		User: &entity.User{
			Id:        userRole.User.Id,
			FirstName: userRole.User.FirstName,
			LastName:  userRole.User.LastName,
			Gender:    helper.ParseGender(userRole.User.Gender),
			BirthDate: userRole.User.BirthDate,
		},
		Role: &entity.Role{
			Id:   userRole.Role.Id,
			Name: userRole.Role.Name,
		},
	}

	// return
	return response, nil
}

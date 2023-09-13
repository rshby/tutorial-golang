package service

import (
	"context"
	"errors"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type RoleService struct {
	RoleRepository *repository.RoleRepository
	Validate       *validator.Validate
}

func NewRoleService(roleRepo *repository.RoleRepository, validate *validator.Validate) *RoleService {
	return &RoleService{
		RoleRepository: roleRepo,
		Validate:       validate,
	}
}

// method insert role
func (r *RoleService) Insert(ctx context.Context, request *web.RequestRoleInsert) (entity.Role, error) {
	wg := &sync.WaitGroup{}
	errChanel := make(chan error, 20)

	// cek data role dengan nama terkait apakah sudah ada di database
	wg.Add(1)
	go func(err chan<- error) {
		logrus.WithField("event", "RoleService-Insert").Info("masuk ke goroutine cek role by name")
		_, errData := r.RoleRepository.GetByName(ctx, request.Name)
		if errData == nil {
			err <- errors.New("record role with name already exist")
			logrus.WithField("event", "RoleServie-Insert").WithField("error", "data role dengan nama tersebut sudah ada").Info("ditemukan error ketika cek role by name")
		} else {
			err <- nil
		}
		wg.Done()
		logrus.WithField("event", "RoleService-Insert").Info("selesai proses goroutine cek role by name")
	}(errChanel)

	// validasi
	wg.Add(1)
	go func(err chan<- error) {
		logrus.WithField("event", "RoleService-Insert").Info("masuk gorotine validasi")
		errData := r.Validate.Struct(*request)

		// kirimkan errData ke channel -> bisa berisi nil atau error_value
		err <- errData

		wg.Done()
		logrus.WithField("event", "RoleService-Insert").Info("selesai proses goroutine validasi")
	}(errChanel)

	wg.Wait()
	close(errChanel)
	for e := range errChanel {
		if e != nil {
			// jika ditemukan error -> return with error
			return entity.Role{}, e
		}
	}

	// buat input
	logrus.WithField("event", "RoleService-Insert").Info("proses create object input")
	role := entity.Role{
		Name: request.Name,
	}

	// jalankan perintah insert yang ada di repository
	logrus.WithField("event", "RoleService-Insert").Info("proses insert role ke database menggunakan repository")
	result, err := r.RoleRepository.General.Insert(ctx, &role)
	if err != nil {
		// jika terdapat error saat insert -> return with error internal server
		logrus.WithField("event", "RoleService-Insert").WithField("error", err.Error()).Info("terdapat error ketika insert data role ke database")
		return entity.Role{}, err
	}

	logrus.WithField("event", "RoleService-Insert").Info("proses get data role by Id setelah Insert")
	// success insert data role to database
	roleData, err := r.RoleRepository.GetById(ctx, result.Id)
	if err != nil {
		// jika terdapat error ketika get data -> return with error data not found
		logrus.WithField("event", "RoleService-Insert").WithField("error", err.Error()).Info("terdapat error ketika get data by Id setelah Insert")
		return entity.Role{}, err
	}

	// return
	logrus.WithField("event", "RoleService-Insert").Info("success return response setelah insert data role ke database success")
	return roleData, nil
}

// method update role
func (r *RoleService) Update(ctx context.Context, request *web.RequestRoleUpdate) (entity.Role, error) {
	wg := &sync.WaitGroup{}
	errChannel := make(chan error, 20)

	// validasi request
	wg.Add(1)
	go func(err chan<- error) {
		errValidate := r.Validate.Struct(*request)

		// kirim errValidate ke channel -> bisa berisi nil atau error_value
		err <- errValidate
		wg.Done()
	}(errChannel)

	// cek data ke database apakah ada
	wg.Add(1)
	go func(err chan<- error) {
		_, errData := r.RoleRepository.General.GetById(ctx, request.Id)

		// kirimkan errData ke channel -> bisa berisi nil atau error_value
		err <- errData
		wg.Done()
	}(errChannel)

	wg.Wait()
	close(errChannel)
	for err := range errChannel {
		if err != nil {
			// return with error
			return entity.Role{}, err
		}
	}

	// buat input
	newRole := entity.Role{
		Id:   request.Id,
		Name: request.Name,
	}

	// jalankan perintah update yang ada di repository
	result, err := r.RoleRepository.General.Update(ctx, newRole.Id, &newRole)
	if err != nil {
		// return with error internal server
		return entity.Role{}, err
	}

	role, err := r.RoleRepository.GetById(ctx, result.Id)
	if err != nil {
		// return with error data not found
		return entity.Role{}, errors.New("record role not found")
	}

	// success update
	// return
	return role, nil
}

// method delete role
func (r *RoleService) Delete(ctx context.Context, id int) (string, error) {
	// cek data di database
	role, err := r.RoleRepository.General.GetById(ctx, id)
	if err != nil {
		// return with error data not found
		return "", errors.New("record role not found")
	}

	// jalankan perintah Delete yang ada di repository
	result, err := r.RoleRepository.General.Delete(ctx, &role)
	if err != nil {
		// return with error internal server
		return "", err
	}

	// success delete data role by Id
	// return
	return result, nil
}

// method get all data roles
func (r *RoleService) GetAll(ctx context.Context) ([]entity.Role, error) {
	// jalankan perintah getall yang ada di repository
	roles, err := r.RoleRepository.GetAll(ctx)
	if err != nil {
		// return with error data not found
		return []entity.Role{}, err
	}

	// success get all data roles
	// return
	return roles, nil
}

// method get role by id
func (r *RoleService) GetById(ctx context.Context, id int) (entity.Role, error) {
	// jalankan perintah getbyid yang ada di reppository
	role, err := r.RoleRepository.GetById(ctx, id)
	if err != nil {
		// return with error data not found
		return entity.Role{}, err
	}

	// success get data by id
	// return
	return role, nil
}

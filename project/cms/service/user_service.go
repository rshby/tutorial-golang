package service

import (
	"cms/model/dto"
	"cms/model/entity"
	"cms/repository"
	"context"
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)

// create object struct user-service
type UserService struct {
	UserRepository *repository.UserRepository
	Validate       *validator.Validate
}

// function provider to create object UserService
func NewUserService(userRepo *repository.UserRepository, validate *validator.Validate) *UserService {
	return &UserService{
		UserRepository: userRepo,
		Validate:       validate,
	}
}

// method Get User by email
func (u *UserService) GetByEmail(ctx context.Context, email string) (*dto.UsersByEmail, error) {
	// call procedure getByEmail in repository
	user, err := u.UserRepository.GetByEmail(ctx, email)

	// jika ada error ketika get data
	if err != nil {
		return nil, err
	}

	gender := ""
	if user.User.Gender == "M" {
		gender = "Male"
	} else {
		gender = "Female"
	}

	user.User.Gender = gender

	// success get data -> return
	return user, nil
}

// method Get All Users
func (u *UserService) GetAll(ctx context.Context) ([]*dto.UserResponse, error) {
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	// call procedure in repository
	result, err := u.UserRepository.GetAll(ctx)

	// jika ada kesalahan ketika get all data
	if err != nil {
		return nil, err
	}

	// create object response
	response := []*dto.UserResponse{}

	for _, data := range result {
		wg.Add(1)
		go func(dt entity.User) {
			defer mtx.Unlock()
			defer wg.Done()

			// lock mutex
			mtx.Lock()

			gender := ""
			if dt.Gender == "M" {
				gender = "Male"
			} else {
				gender = "Female"
			}

			// create object from dto
			item := dto.UserResponse{
				Id:        dt.ID,
				FirstName: dt.FirstName,
				LastName:  dt.LastName,
				FullName:  fmt.Sprintf("%v %v", dt.FirstName, dt.LastName),
				Gender:    gender,
				Address:   dt.Address,
				CreatedAt: dt.CreatedAt.Format("2006-01-02 15:04:05"),
			}

			// append item to response
			response = append(response, &item)
		}(data)
	}

	// wait all thread done
	wg.Wait()

	// return response
	return response, nil
}

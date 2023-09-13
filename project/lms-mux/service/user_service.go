package service

import (
	"context"
	"fmt"
	"lms-mux/helper"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/repository"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	UserRepository       *repository.UserRepository
	AccountRepository    *repository.AccountRepository
	TakenClassRepository *repository.TakenClassRepository
	UserRoleRepository   *repository.UserRoleRepository
	Validate             *validator.Validate
}

func NewUserService(userRepo *repository.UserRepository, accountRepo *repository.AccountRepository, takenclassRepo *repository.TakenClassRepository, userRoleRepo *repository.UserRoleRepository, validate *validator.Validate) *UserService {
	return &UserService{
		UserRepository:       userRepo,
		AccountRepository:    accountRepo,
		TakenClassRepository: takenclassRepo,
		UserRoleRepository:   userRoleRepo,
		Validate:             validate,
	}
}

// method Insert user
func (u *UserService) Insert(ctx context.Context, request *web.RequestUserInsert) (web.ResponseUser, error) {
	// validasi request
	err := u.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return web.ResponseUser{}, err
	}

	// buat input
	user := entity.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Gender:      request.Gender,
		BirthDate:   request.BirthDate,
		AddressId:   request.AddressId,
		EducationId: request.EducationId,
	}

	// jalankan perintah Insert yang ada di repository
	resultInsert, err := u.UserRepository.General.Insert(ctx, &user)
	if err != nil {
		// return with error internal server
		return web.ResponseUser{}, err
	}

	// success insert data user
	result, err := u.GetUserInformationById(ctx, resultInsert.Id)
	if err != nil {
		// return with error data not found
		return web.ResponseUser{}, err
	}

	// return
	return result, nil
}

// method update data user by id
func (u *UserService) Update(ctx context.Context, request *web.RequestUserUpdate) (web.ResponseUser, error) {
	// cek id apakah data yang akan diupdate ada di database
	oldUser, err := u.UserRepository.General.GetById(ctx, request.Id)
	if err != nil {
		// return with error not found
		return web.ResponseUser{}, err
	}

	// validasi request body
	err = u.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return web.ResponseUser{}, err
	}

	// buat inputan
	user := entity.User{
		Id:          oldUser.Id,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Gender:      request.Gender,
		AddressId:   request.AddressId,
		EducationId: request.EducationId,
	}

	// jalankan perintah update yang ada di repository
	result, err := u.UserRepository.General.Update(ctx, request.Id, &user)
	if err != nil {
		// return with error internal server
		return web.ResponseUser{}, err
	}

	// success update data user by Id
	response, err := u.GetUserInformationById(ctx, result.Id)
	if err != nil {
		// return with error not found
		return web.ResponseUser{}, err
	}

	// return
	return response, nil
}

// method Delete user by Id
func (u *UserService) Delete(ctx context.Context, userId int) (string, error) {
	// cek data user_id apakah ada di database
	oldUser, err := u.UserRepository.General.GetById(ctx, userId)
	if err != nil {
		// return with error data not found
		return "", err
	}

	wg := &sync.WaitGroup{}
	errChan := make(chan error, 10)

	// jalankan perintah delete yang ada di user repository
	// 1. get account by user_id
	wg.Add(1)
	go func(eChan chan error) {
		fmt.Println("masuk goroutine account")
		account, err := u.AccountRepository.GetAccountByUserId(ctx, oldUser.Id)

		// jika ada data account
		if err == nil {
			fmt.Println("ada data account")

			// hapus account by user_id
			_, err = u.AccountRepository.General.Delete(ctx, &account)
			eChan <- err
		} else {
			// tidak ada data account -> langsung kirim error nil
			eChan <- nil
		}
		fmt.Println("keluar goroutine account")
		wg.Done()
	}(errChan)

	//2. get all data takenclasses by user_Id
	wg.Add(1)
	go func(eChan chan error) {
		fmt.Println("masuk goroutine takenclass")
		takenClasses, err := u.TakenClassRepository.GetTakenClassByuserId(ctx, oldUser.Id)

		// jika ada data takenclass
		if err == nil {
			fmt.Println("ada data takenclass")

			// hapus seluruh data takenclass dengan user_id terkait
			for _, takenClass := range takenClasses {
				wg.Add(1)
				go func(eChannel chan error, tc *entity.TakenClass) {
					_, err := u.TakenClassRepository.General.Delete(ctx, tc)
					eChannel <- err
					wg.Done()
				}(errChan, &takenClass)
			}
		} else {
			// tidak ada data takenclass -> langsung kirim error nil
			eChan <- nil
		}
		fmt.Println("keluar goroutine takenclass")
		wg.Done()
	}(errChan)

	// 3. get all data userroles by user_id
	wg.Add(1)
	go func(eChan chan error) {
		fmt.Println("masuk goroutine userRole")
		userRoles, err := u.UserRoleRepository.GetUserRolesByUserId(ctx, oldUser.Id)

		// jika ada data userRoles
		if err == nil {
			fmt.Println("ada data userRole")

			// hapus semua data userRoles dengan user_id terkait
			for _, userRole := range userRoles {
				_, err := u.UserRoleRepository.General.Delete(ctx, &userRole)
				eChan <- err
			}
		} else {
			// jika tidak ada data userrole -> langsung kirim error nil
			eChan <- nil
		}

		fmt.Println("keluar goroutine userRole")
		wg.Done()
	}(errChan)

	wg.Wait()
	close(errChan)

	fmt.Println("errChan dipindah ke errMessage")
	for errMessgae := range errChan {
		// jika ada error -> return error
		if errMessgae != nil {
			return "", errMessgae
		}
	}

	// 4. hapus user by id
	_, err = u.UserRepository.General.Delete(ctx, &oldUser)
	if err != nil {
		// return error with internal server
		return "", err
	}

	// success delete data user
	// return
	return "success delete data user", nil
}

// method get full information User
func (u *UserService) GetUserInformationById(ctx context.Context, userId int) (web.ResponseUser, error) {
	// jalankan perintah GetUserInformationById yang ada di repository
	result, err := u.UserRepository.GetUserInformationById(ctx, userId)
	if err != nil {
		// return with error data not found
		return web.ResponseUser{}, err
	}

	// mapping gender
	var gender string
	if result.Gender == "L" {
		gender = "Laki-Laki"
	} else {
		gender = "Perempuan"
	}

	// success get user information by Id
	user := web.ResponseUser{
		Id:        result.Id,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Gender:    gender,
		BirthDate: result.BirthDate.String(),
		Account:   result.Account,
		Address: &web.ResponseFullAddress{
			Id:          result.Address.Id,
			Street:      result.Address.Street,
			SubDistrict: result.Address.SubDistrict.Name,
			ZipCode:     result.Address.SubDistrict.ZipCode,
			District:    result.Address.SubDistrict.District.Name,
			City:        result.Address.SubDistrict.District.City.Name,
			Province:    result.Address.SubDistrict.District.City.Province.Name,
		},
		Education: &web.ResponseEducation{
			Id:         result.Education.Id,
			Major:      result.Education.Major,
			Level:      result.Education.Level,
			University: result.Education.University.Name,
		},
	}

	// return
	return user, nil
}

// method GetById
func (u *UserService) GetById(ctx context.Context, id int) (entity.User, error) {
	// jalankkan getbyid
	user, err := u.UserRepository.GetUserInformationById(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	// success
	// return
	return user, nil
}

// method get all users information
func (u *UserService) GetAll(ctx context.Context) ([]web.ResponseUser, error) {
	// jalankan perintah getall users yang ada di repository
	results, err := u.UserRepository.GetAllUsersInformation(ctx)
	if err != nil {
		// return with error data not found
		return []web.ResponseUser{}, err
	}

	// success get all users data
	var users []web.ResponseUser
	wg := &sync.WaitGroup{}
	mtx := &sync.RWMutex{}

	for i := range results {
		wg.Add(1)
		go func(user *entity.User) {
			mtx.Lock()

			// append ke variabel users
			users = append(users, web.ResponseUser{
				Id:        user.Id,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Gender:    helper.ParseGender(user.Gender),
				BirthDate: user.BirthDate.String(),
				Account:   user.Account,
				Address: &web.ResponseFullAddress{
					Id:          user.Address.Id,
					Street:      user.Address.Street,
					SubDistrict: user.Address.SubDistrict.Name,
					ZipCode:     user.Address.SubDistrict.ZipCode,
					District:    user.Address.SubDistrict.District.Name,
					City:        user.Address.SubDistrict.District.City.Name,
					Province:    user.Address.SubDistrict.District.City.Province.Name,
				},
				Education: &web.ResponseEducation{
					Id:         user.Education.Id,
					Major:      user.Education.Major,
					Level:      user.Education.Level,
					University: user.Education.University.Name,
				},
			})

			mtx.Unlock()
			wg.Done()
		}(&results[i])
	}

	// return
	wg.Wait()
	return users, nil
}

// method get by Id

// method get all user by batch
func (u *UserService) GetAllByBatch(ctx context.Context) ([]web.ResponseUser, error) {
	var totalCount int64
	var response []web.ResponseUser
	wgQuery := &sync.WaitGroup{}
	wgLogic := &sync.WaitGroup{}
	mtx := &sync.RWMutex{}

	u.UserRepository.DB.WithContext(ctx).Model(&entity.User{}).Count(&totalCount)
	logrus.Info(totalCount)

	usersChannel := make(chan []entity.User)

	go func(chanData chan<- []entity.User, total int64) {
		batch := 12000

		for i := 1; i <= int(total); i += int(batch) {
			endId := i + int(batch)

			wgQuery.Add(1)
			go func(chanData chan<- []entity.User, beginId int, endId int) {
				defer wgQuery.Done()

				datas := []entity.User{}
				u.UserRepository.DB.WithContext(ctx).Model(&entity.User{}).Preload("Address.SubDistrict.District.City.Province").Preload("Account").Preload("Education.University").Where("users.id>=? AND users.id<?", beginId, endId).Find(&datas)
				chanData <- datas // send data to channel
			}(chanData, i, endId)
		}

		wgQuery.Wait()
		close(chanData) // close channel
	}(usersChannel, totalCount)

	// consume data from channel
	for users := range usersChannel { // ada 10 -> 1.. 2.., 3.. wadas(1,2,3, 100, 101, 102, ) (100, 1, 800, )
		wgLogic.Add(1)
		go func(users []entity.User) {
			defer wgLogic.Done()
			mtx.Lock()

			for _, user := range users {

				// append ke variabel response
				response = append(response, web.ResponseUser{
					Id:        user.Id,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Gender:    helper.ParseGender(user.Gender),
					BirthDate: user.BirthDate.String(),
					Account:   user.Account,
					Address: &web.ResponseFullAddress{
						Id:          user.Address.Id,
						Street:      user.Address.Street,
						SubDistrict: user.Address.SubDistrict.Name,
						ZipCode:     user.Address.SubDistrict.ZipCode,
						District:    user.Address.SubDistrict.District.Name,
						City:        user.Address.SubDistrict.District.City.Name,
						Province:    user.Address.SubDistrict.District.City.Province.Name,
					},
					Education: &web.ResponseEducation{
						Id:         user.Education.Id,
						Major:      user.Education.Major,
						Level:      user.Education.Level,
						University: user.Education.University.Name,
					},
				})
			}
			mtx.Unlock()
		}(users)
	}

	wgLogic.Wait()
	return response, nil
}

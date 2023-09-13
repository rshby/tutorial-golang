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

type AccountService struct {
	AccountRepository *repository.AccountRepository
	Validate          *validator.Validate
}

func NewAccountService(accountRepo *repository.AccountRepository, validate *validator.Validate) *AccountService {
	return &AccountService{
		AccountRepository: accountRepo,
		Validate:          validate,
	}
}

// method insert
func (a *AccountService) Insert(ctx context.Context, request *web.RequestAccountInsert) (web.ResponseAccount, error) {
	// cek email apakah sudah ada di database
	_, err := a.AccountRepository.GetAccountByEmail(ctx, request.Email)

	// jika data account dengan email tersebut sudah ada
	if err == nil {
		// return wit error data already exist
		return web.ResponseAccount{}, errors.New("data with email already exist")
	}

	// validasi request
	err = a.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> gagal validasi
		return web.ResponseAccount{}, err
	}

	// buat inputan
	hashedPassword := helper.HashedPassword(request.Password)
	account := entity.Account{
		Email:    request.Email,
		Password: hashedPassword,
		UserId:   request.UserId,
	}

	// jalankan perintah insert yang ada di repository general
	accountInsert, err := a.AccountRepository.General.Insert(ctx, &account)
	if err != nil {
		// return with error internal server
		return web.ResponseAccount{}, err
	}

	// success insert -> get data by Id
	result, err := a.AccountRepository.GetById(ctx, accountInsert.Id)
	if err != nil {
		// return with error data not found
		return web.ResponseAccount{}, err
	}

	response := web.ResponseAccount{
		Id:       result.Id,
		Email:    result.Email,
		Password: result.Password,
		User: &web.ResponseUser{
			FirstName: result.User.FirstName,
			LastName:  result.User.LastName,
			Gender:    helper.ParseGender(result.User.Gender),
			BirthDate: result.User.BirthDate.String(),
			Education: &web.ResponseEducation{
				Major:      result.User.Education.Major,
				Level:      result.User.Education.Level,
				University: result.User.Education.University.Name,
			},
		},
	}

	// return response
	return response, nil
}

// method Update account
func (a *AccountService) Update(ctx context.Context, request *web.RequestUpdateAccount) (web.ResponseAccount, error) {
	// cek apakah id ada di database
	wg := &sync.WaitGroup{}
	errChannel := make(chan error, 50)

	// cek data dengan id tersebut apakah ada di database
	wg.Add(1)
	go func(err chan error) {
		logrus.WithField("event", "AccountService-Update").Info("masuk ke goroutine cek id data ke database")
		_, errData := a.AccountRepository.General.GetById(ctx, request.Id)

		// kirim errData ke channel -> bisa berisi nil atau error_value
		err <- errData

		wg.Done()
	}(errChannel)

	// validasi request
	wg.Add(1)
	go func(err chan<- error) {
		logrus.WithField("event", "AccountService-Update").Info("masuk ke goroutine cek validasi")
		errValidate := a.Validate.Struct(*request)

		// kirim errValidate ke channel -> bisa berisi nil atau error_value
		err <- errValidate

		wg.Done()
	}(errChannel)

	wg.Wait()
	close(errChannel)
	for err := range errChannel {
		if err != nil {
			logrus.WithField("event", "AccountService-Update").WithField("error", err.Error()).Info("Ketemu nih errornya")
			// jika ada error_value pada channel -> return with error
			return web.ResponseAccount{}, err
		}
	}

	// buat input
	account := entity.Account{
		Id:       request.Id,
		Email:    request.Email,
		Password: helper.HashedPassword(request.Password),
		UserId:   request.UserId,
	}

	// jalankan perintah Update yang ada di repository
	result, err := a.AccountRepository.General.Update(ctx, account.Id, &account)
	if err != nil {
		// jika gagal update -> return with error internal server
		logrus.WithField("event", "AccountService-Update").WithField("error", err.Error()).Info("gagal update")
		return web.ResponseAccount{}, err
	}
	logrus.WithField("event", "AccountService-Update").Info("success update data account by Id")

	// success update -> get data by Id
	data, err := a.AccountRepository.GetAccountByUserId(ctx, result.UserId)
	if err != nil {
		// jika data account tidak ada -> return with error not found
		return web.ResponseAccount{}, errors.New("record account not found")
	}

	response := web.ResponseAccount{
		Id:       data.Id,
		Email:    data.Email,
		Password: data.Password,
		User: &web.ResponseUser{
			FirstName: data.User.FirstName,
			LastName:  data.User.LastName,
			Gender:    helper.ParseGender(data.User.Gender),
			BirthDate: data.User.BirthDate.String(),
			Education: &web.ResponseEducation{
				Major:      data.User.Education.Major,
				Level:      data.User.Education.Level,
				University: data.User.Education.University.Name,
			},
		},
	}

	// return
	logrus.WithField("event", "AccountService-Update").Info("return response after success update account")
	return response, nil
}

// method delete account
func (a *AccountService) Delete(ctx context.Context, id int) (string, error) {
	// ambil data account by Id
	logrus.WithField("event", "AccountService-Delete").Info("cek data account dengan Id yang akan dihapus")
	account, err := a.AccountRepository.GetById(ctx, id)
	if err != nil {
		// jika data tidak ada -> return with error data not found
		logrus.WithField("event", "AccountService-Delete").WithField("error", err.Error()).Info("ditemukan error sewaktu cek data account")
		return "", errors.New("record account not found")
	}

	// jalankan perintah delete yang ada di repository general
	logrus.WithField("event", "AccountService-Delete").Info("proses delete account by Id")
	_, err = a.AccountRepository.General.Delete(ctx, &account)
	if err != nil {
		// jika gagal delete -> return with error internal server
		logrus.WithField("event", "AccountService-Delete").WithField("error", err.Error()).Info("ditemukan error ketika proses delete account by Id")
		return "", err
	}

	// success delete data account
	// return
	logrus.WithField("event", "AccountService-Delete").Info("success delete account by Id")
	return "success delete data account by Id", nil
}

// method Get Data by Id
func (a *AccountService) GetById(ctx context.Context, id int) (web.ResponseAccount, error) {
	logrus.WithField("event", "AccountService-GetById").Info("proses get data account by Id")
	result, err := a.AccountRepository.GetById(ctx, id)
	if err != nil {
		// jika tidak ada datanya -> return with error not found
		logrus.WithField("event", "AccountService-GetById").WithField("error", err.Error()).Info("ditemukan error pada proses get data account by Id")
		return web.ResponseAccount{}, errors.New("record account not found")
	}

	// success get data by Id
	response := web.ResponseAccount{
		Id:       result.Id,
		Email:    result.Email,
		Password: result.Password,
		User: &web.ResponseUser{
			FirstName: result.User.FirstName,
			LastName:  result.User.LastName,
			Gender:    helper.ParseGender(result.User.Gender),
			BirthDate: result.User.BirthDate.String(),
			Education: &web.ResponseEducation{
				Major:      result.User.Education.Major,
				Level:      result.User.Education.Level,
				University: result.User.Education.University.Name,
			},
		},
	}
	logrus.WithField("event", "AccountService-GetById").Info("success create object response")

	// success get data account by Id
	// return
	logrus.WithField("event", "AccountService-GetById").Info("success return response")
	return response, nil
}

// method get all data account
func (a *AccountService) GetAll(ctx context.Context) ([]web.ResponseAccount, error) {
	var accounts []web.ResponseAccount

	logrus.WithField("event", "AccountService-GetAll").Info("proses get all data accounts from repository")
	results, err := a.AccountRepository.GetAll(ctx)
	if err != nil {
		// jika tidak ada datanya -> return with error not found
		logrus.WithField("event", "AccountService-GetAll").WithField("error", err.Error()).Info("ditemukan error ketika proses get all data accounts")
		return []web.ResponseAccount{}, err
	}

	// success get all data accounts
	for _, result := range results {
		accounts = append(accounts, web.ResponseAccount{
			Id:       result.Id,
			Email:    result.Email,
			Password: result.Password,
			User: &web.ResponseUser{
				FirstName: result.User.FirstName,
				LastName:  result.User.LastName,
				Gender:    helper.ParseGender(result.User.Gender),
				BirthDate: result.User.BirthDate.String(),
				Education: &web.ResponseEducation{
					Major:      result.User.Education.Major,
					Level:      result.User.Education.Level,
					University: result.User.Education.University.Name,
				},
			},
		})
	}
	logrus.WithField("event", "AccountService-GetAll").Info("success append data response all accounts")

	// return
	logrus.WithField("event", "AccountService-GetAll").Info("success return response")
	return accounts, nil
}

// method get data account by email
func (a *AccountService) GetAccountByEmail(ctx context.Context, email string) (web.ResponseAccount, error) {
	logrus.WithField("event", "AccountService-GetAccountByEmail").Info("proses get data account by email")
	cekData, err := a.AccountRepository.GetAccountByEmail(ctx, email)
	if err != nil {
		// jika tidak ada datanya -> return with error not found
		logrus.WithField("event", "AccountService-GetAccountByEmail").WithField("error", err.Error()).Info("ditemukan error ketika proses get data account by email")
		return web.ResponseAccount{}, errors.New("record account not found")
	}

	// success get data
	logrus.WithField("event", "AccountService-GetAccountByEmail").Info("proses get data lengkap account by Id")
	account, err := a.GetById(ctx, cekData.Id)
	if err != nil {
		// jika tidak ada datanya -> return with error not found
		logrus.WithField("event", "AccountService-GetAccountByEmail").WithField("error", err.Error()).Info("ditemukan error ketika proses get data account by Id")
		return web.ResponseAccount{}, errors.New("record account not found")
	}

	// success get data lengkap
	// return
	logrus.WithField("event", "AccountService-GetAccountByEmail").Info("success return response account")
	return account, nil
}

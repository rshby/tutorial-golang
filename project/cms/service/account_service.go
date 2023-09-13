package service

import (
	"cms/helper"
	"cms/model/auth"
	"cms/model/dto"
	"cms/model/entity"
	"cms/repository"
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

// create object account service
type AccountService struct {
	AccountRepository *repository.AccountRepository
	UserRepository    *repository.UserRepository
	Validate          *validator.Validate
}

// function provider to create new AccountService object
func NewAccountService(accRepo *repository.AccountRepository, userRepo *repository.UserRepository, validate *validator.Validate) *AccountService {
	return &AccountService{
		AccountRepository: accRepo,
		UserRepository:    userRepo,
		Validate:          validate,
	}
}

// method Create Account
func (a *AccountService) CreateAccount(ctx context.Context, request *dto.CreateAccountRequest) (*dto.CreateAccountResponse, error) {
	// cek required validasi request
	err := a.Validate.Struct(*request)
	if err != nil {
		// return with error bad request -> kena validasi required
		return nil, err
	}

	// cek jika data dengan email yang sama sudah ada
	_, err = a.AccountRepository.GetByEmail(ctx, request.Email)
	if err == nil {
		// return data dengan email tersebut sudah ada di database
		return nil, errors.New("data with same email already exist")
	}

	// cek jika data dengan username yang sama sudah ada di database
	_, err = a.AccountRepository.GetByUsername(ctx, request.Username)
	if err == nil {
		return nil, errors.New("data with same username already exist in database")
	}

	// buat inputan untuk create user terlebih dahulu
	userInput := entity.User{
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		IdentityId: request.IdentityId,
		Gender:     request.Gender,
		Address:    request.Address,
	}

	// proses insert user
	userInsert, err := a.UserRepository.Insert(ctx, &userInput)
	if err != nil {
		// error ketika insert user
		return nil, err
	}

	// hased password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	// error ketika hashed passoword
	if err != nil {
		errDeleteUser := a.UserRepository.Delete(ctx, userInsert.ID)
		if errDeleteUser != nil {
			return nil, errDeleteUser
		}

		return nil, err
	}

	// buat inputan untuk insert account
	accountInput := entity.Account{
		Email:    request.Email,
		Username: request.Username,
		Password: string(hashedPassword),
		UserId:   userInsert.ID,
	}

	// proses insert account
	accountInsert, err := a.AccountRepository.Insert(ctx, &accountInput)

	// error ketika insert account
	if err != nil {
		errDeleteUser := a.UserRepository.Delete(ctx, userInsert.ID)
		if errDeleteUser != nil {
			return nil, errDeleteUser
		}

		return nil, err
	}

	// success insert -> mapping response
	userDetail, err := a.UserRepository.GetById(ctx, userInsert.ID)

	// error ketika get user yang diinsert
	if err != nil {
		errDeleteUser := a.UserRepository.Delete(ctx, userInsert.ID)
		if errDeleteUser != nil {
			return nil, errDeleteUser
		}

		errDeleteAccount := a.AccountRepository.Delete(ctx, accountInsert.ID)
		if errDeleteAccount != nil {
			return nil, errDeleteAccount
		}

		return nil, err
	}

	accountDetail, err := a.AccountRepository.GetById(ctx, accountInsert.ID)

	// error ketika get account yang diinsert
	if err != nil {
		errDeleteUser := a.UserRepository.Delete(ctx, userInsert.ID)
		if errDeleteUser != nil {
			return nil, errDeleteUser
		}

		errDeleteAccount := a.AccountRepository.Delete(ctx, accountInsert.ID)
		if errDeleteAccount != nil {
			return nil, errDeleteAccount
		}

		return nil, err
	}

	// mapping gender
	var gender = ""
	if userDetail.Gender == "M" {
		gender = "Laki-Laki"
	} else {
		gender = "Perempuan"
	}

	// mapping object response
	response := dto.CreateAccountResponse{
		ID_User:    userDetail.ID,
		ID_Account: accountDetail.ID,
		AccountDetail: &dto.AccountDetail{
			Id:        accountDetail.ID,
			Email:     accountDetail.Email,
			Username:  accountDetail.Username,
			Password:  accountDetail.Password,
			CreatedAt: accountDetail.CreatedAt.Format("2006-01-02 15:04:05"),
		},
		UserDetail: &entity.User{
			ID:         userDetail.ID,
			FirstName:  userDetail.FirstName,
			LastName:   userDetail.LastName,
			IdentityId: userDetail.IdentityId,
			Gender:     gender,
			Address:    userDetail.Address,
			CreatedAt:  userDetail.CreatedAt,
		},
	}

	// success mapping -> return response
	return &response, nil
}

// method check account by email
func (a *AccountService) CheckAccountEmail(ctx context.Context, request *dto.CheckAccountEmailRequest) (*dto.CheckAccountResponse, error) {
	// validate request
	err := a.Validate.Struct(*request)

	// jika ada error gagal validasi required
	if err != nil {
		return nil, err
	}

	// call procedure getAccountByEmail
	account, err := a.AccountRepository.GetByEmail(ctx, request.Email)

	// jika terdapat error ketika proses get data by email
	if err != nil {
		return nil, err
	}

	// success get data -> mapping to response
	response := dto.CheckAccountResponse{
		AlreadyExist: true,
		Account: &dto.AccountDetail{
			Id:        account.ID,
			Email:     account.Email,
			Username:  account.Username,
			Password:  account.Password,
			CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	// return
	return &response, nil
}

// method check account by username
func (a *AccountService) CheckAccountUsername(ctx context.Context, request *dto.CheckAccountUsernameRequest) (*dto.CheckAccountResponse, error) {
	// validasi required
	err := a.Validate.Struct(*request)

	// jika gagal validasi required
	if err != nil {
		return nil, err
	}

	// call procedure get by username in repository
	account, err := a.AccountRepository.GetByUsername(ctx, request.Username)

	// jika ada error ketika get data by username
	if err != nil {
		return nil, err
	}

	// success get data -> mapping to response
	response := dto.CheckAccountResponse{
		AlreadyExist: true,
		Account: &dto.AccountDetail{
			Id:        account.ID,
			Email:     account.Email,
			Username:  account.Username,
			Password:  account.Password,
			CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	// return
	return &response, nil
}

// method login
func (a *AccountService) Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error) {
	// validasi required
	err := a.Validate.Struct(*request)

	// jika gagal validasi required
	if err != nil {
		return nil, err
	}

	// cek username apakah ada di database
	var dataLogin *entity.Account
	dataLogin, err = a.AccountRepository.GetByUsername(ctx, request.Username)
	if err != nil {
		// jika error -> coba cek dengan email apabila user login menggunakan email
		cekEmail, errCheck := a.AccountRepository.GetByEmail(ctx, request.Username)

		// jika dicek berdasarkan email juga tetap tidak ketemu -> return error not found
		if errCheck != nil {
			return nil, errors.New("username or email not exist in database")
		} else {
			// jika ketemu datanya
			dataLogin = cekEmail
		}
	}

	// cocokan dengan password
	err = bcrypt.CompareHashAndPassword([]byte(dataLogin.Password), []byte(request.Password))

	// jika password salah
	if err != nil {
		return nil, errors.New("wrong password")
	}

	// lolos cek username dan password -> buat token jwt
	claim := auth.JWTClaim{
		Email: dataLogin.Email,
		ClaimRegistered: jwt.RegisteredClaims{
			Issuer:    "cms",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	// create token algorithm
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)

	// create token string
	token, err := tokenAlgo.SignedString(auth.JWT_SECRET_KEY)
	if err != nil {
		return nil, err
	}

	// create response login
	response := dto.LoginResponse{
		IsLoggin: true,
		Token:    token,
	}

	return &response, nil
}

// method request OTP
func (a *AccountService) RequestOTP(ctx context.Context, request *dto.OtpRequest) (*dto.OtpResponse, error) {
	// validasi request required
	err := a.Validate.Struct(*request)

	// jika gagal validasi required
	if err != nil {
		return nil, err
	}

	// cek email apakah ada di database
	cekEmail, err := a.CheckAccountEmail(ctx, &dto.CheckAccountEmailRequest{Email: request.Email})
	if err != nil {
		return nil, err
	}

	// receive otp number
	otp, err := helper.GenerateOTP(6)
	if err != nil {
		return nil, err
	}

	// create inputan to update otp
	account := entity.Account{
		ID:    cekEmail.Account.Id,
		Email: cekEmail.Account.Email,
		OTP: sql.NullString{
			String: otp,
			Valid:  true,
		},
		ExpiredOTP: sql.NullTime{
			Time:  time.Now().Local().Add(3 * time.Minute),
			Valid: true,
		},
	}

	// call procedure in repository to update otp
	err = a.AccountRepository.UpdateOTP(ctx, &account)

	// jika ada error ketika proses update otp
	if err != nil {
		return nil, err
	}

	// success update otp -> create response
	response := dto.OtpResponse{
		OTP:        account.OTP.String,
		ExpiredOtp: account.ExpiredOTP.Time.Format("2006-01-02 15:04:05"),
	}

	return &response, nil
}

// method Forgot Password
func (a *AccountService) ForgotPassword(ctx context.Context, request *dto.ForgotPasswordRequest) (*dto.ForgotPasswordResponse, error) {
	// validasi request
	err := a.Validate.Struct(*request)

	// jika ada error gagal validasi required
	if err != nil {
		return nil, err
	}

	// get data by email
	accByEmail, err := a.AccountRepository.GetByEmail(ctx, request.Email)

	// jika email tidak ada di database
	if err != nil {
		return nil, errors.New("email not exist in database")
	}

	// cocokan dengan password
	err = bcrypt.CompareHashAndPassword([]byte(accByEmail.Password), []byte(request.OldPassword))

	// jika ada error salah password
	if err != nil {
		return nil, errors.New("old password wrong")
	}

	// cocokan otp
	if accByEmail.OTP.String != request.OTP {
		// jika otp salah
		return nil, errors.New("wrong OTP")
	}

	// cek expired otp yang ada
	if time.Now().After(accByEmail.ExpiredOTP.Time) {
		// jika otp sudah expired
		return nil, errors.New("OTP expired")
	}

	// cek new password & confirm new password
	if request.NewPassword != request.ConfirmNewPassword {
		// jika beda
		return nil, errors.New("new password and confirm password mus be same")
	}

	// ---- lolos semua validasi ---
	// update password, otp, and expired OTP
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// update password
	err = a.AccountRepository.UpdatePassword(ctx, accByEmail.ID, string(hashedNewPassword))
	if err != nil {
		return nil, err
	}

	// update otp
	err = a.AccountRepository.UpdateOTP(ctx, &entity.Account{
		ID: accByEmail.ID,
		OTP: sql.NullString{
			Valid: false,
		},
		ExpiredOTP: sql.NullTime{
			Valid: false,
		},
	})

	// jika ada error update otp
	if err != nil {
		return nil, err
	}

	// --- success update password dan otp ---
	// get data by id
	account, err := a.AccountRepository.GetById(ctx, accByEmail.ID)
	if err != nil {
		return nil, err
	}

	// mapping to response dto.ForgotPasswordResponse
	response := dto.ForgotPasswordResponse{
		Email:    account.Email,
		Password: accByEmail.Password,
	}

	// return response
	return &response, nil
}

// method ChangePassword
func (a *AccountService) ChangePassword(ctx context.Context, request *dto.ChangePasswordRequest) (*dto.ChangePasswordResponse, error) {
	// validasi request
	err := a.Validate.Struct(*request)

	// jika ada error gagal validasi
	if err != nil {
		return nil, err
	}

	// cek data by email -> apakah email ada di database
	accountByEmail, err := a.AccountRepository.GetByEmail(ctx, request.Email)

	// jika data tidak ditemukan
	if err != nil {
		return nil, errors.New("data with email not exist in database")
	}

	// cocokan old password
	err = bcrypt.CompareHashAndPassword([]byte(accountByEmail.Password), []byte(request.OldPassword))

	// jika old password tidak sama
	if err != nil {
		return nil, errors.New("wrong old password")
	}

	// cocokan new_password dan confirm_password
	if request.NewPassword != request.ConfirmPassword {
		return nil, errors.New("new password and confirm password must be same")
	}

	// --- lolos semua validasi ---
	// create hasehd password from new_password
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)

	// jika gagal hasehd password
	if err != nil {
		return nil, err
	}

	// call procedure update_password in repository
	err = a.AccountRepository.UpdatePassword(ctx, accountByEmail.ID, string(newHashedPassword))

	// jika ada kesalahan update password
	if err != nil {
		return nil, err
	}

	// success update password -> get data by id
	account, err := a.AccountRepository.GetById(ctx, accountByEmail.ID)
	if err != nil {
		return nil, err
	}

	// create response
	response := dto.ChangePasswordResponse{
		Email:    account.Email,
		Password: account.Password,
	}

	// return
	return &response, nil
}

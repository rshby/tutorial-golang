package repository

import (
	"context"
	"errors"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB      *gorm.DB
	General *GeneralRepository[entity.Account]
}

func NewAccountRepository(db *gorm.DB, general *GeneralRepository[entity.Account]) *AccountRepository {
	return &AccountRepository{
		DB:      db,
		General: general,
	}
}

// method get account by Id
func (a *AccountRepository) GetById(ctx context.Context, id int) (entity.Account, error) {
	var account entity.Account

	result := a.DB.WithContext(ctx).Model(&entity.Account{}).Preload("User.Education.University").Where("accounts.id=?", id).First(&account)
	if result.Error != nil {
		// return with error data not found
		return entity.Account{}, result.Error
	}

	// success get data account by Id
	// return
	return account, nil
}

// method get account by user_id
func (a *AccountRepository) GetAccountByUserId(ctx context.Context, userId int) (entity.Account, error) {
	var account entity.Account

	result := a.DB.WithContext(ctx).Model(&entity.Account{}).Preload("User.Education.University").Where("accounts.user_id=?", userId).First(&account)
	if result.RowsAffected == 0 {
		// return with error data not found
		return entity.Account{}, errors.New("record account not found")
	}

	// success get data account by user_id
	// return
	return account, nil
}

// method get account by email
func (a *AccountRepository) GetAccountByEmail(ctx context.Context, email string) (entity.Account, error) {
	var account entity.Account

	result := a.DB.WithContext(ctx).Model(&entity.Account{}).Preload("User.Education.University").Where("accounts.email=?", email).First(&account)
	if result.Error != nil {
		// return with error data not found
		return entity.Account{}, result.Error
	}

	// success get data account by email
	// return
	return account, nil
}

// method get all
func (a *AccountRepository) GetAll(ctx context.Context) ([]entity.Account, error) {
	var accounts []entity.Account

	result := a.DB.WithContext(ctx).Model(&entity.Account{}).Preload("User.Education.University").Find(&accounts)
	if result.RowsAffected == 0 {
		// jika tidak ada datanya -> return with error data not found
		return []entity.Account{}, errors.New("record accounts not found")
	}

	// success get all data accounts
	// return
	return accounts, nil
}

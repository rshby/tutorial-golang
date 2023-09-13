package repository

import (
	"cms/model/entity"
	"context"
	"database/sql"
	"errors"
)

// create object accountRepository
type AccountRepository struct {
	DB     *sql.DB
	Entity *entity.Account
}

// create function provider
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB:     db,
		Entity: &entity.Account{},
	}
}

// insert data Account
func (a *AccountRepository) Insert(ctx context.Context, entity *entity.Account) (*entity.Account, error) {
	query := "INSERT INTO accounts (email, username, password, user_id) VALUES (?, ?, ?, ?)"
	result, err := a.DB.ExecContext(ctx, query, entity.Email, entity.Username, entity.Password, entity.UserId)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	entity.ID = id

	// success -> return object
	return entity, nil
}

// method Get By ID
func (a *AccountRepository) GetById(ctx context.Context, id int64) (*entity.Account, error) {
	query := "SELECT id, email, username, password, otp, expired_otp, created_at, user_id FROM accounts WHERE id=? LIMIT 1"
	row, err := a.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create object intitiate
	account := entity.Account{}

	if row.Next() {
		// scan hasil query ke object user
		err := row.Scan(&account.ID, &account.Email, &account.Username, &account.Password, &account.OTP, &account.ExpiredOTP, &account.CreatedAt, &account.UserId)

		// jika ada error ketika proses scan
		if err != nil {
			return nil, err
		}

		// jika tidak ada error ketika scan -> success scan
		return &account, nil
	} else {
		// jika tidak ada datanya
		return nil, errors.New("record not found")
	}
}

// method GetAccount by Email
func (a *AccountRepository) GetByEmail(ctx context.Context, email string) (*entity.Account, error) {
	query := "SELECT id, email, username, password, otp, expired_otp, created_at, user_id FROM accounts WHERE email=? LIMIT 1"
	row, err := a.DB.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create object initiate
	account := entity.Account{}

	if row.Next() {
		// scan hasil query ke object account
		err := row.Scan(&account.ID, &account.Email, &account.Username, &account.Password, &account.OTP, &account.ExpiredOTP, &account.CreatedAt, &account.UserId)

		// jika ada error saat proses Scan
		if err != nil {
			return nil, err
		}

		// success Scan -> return response
		return &account, nil
	} else {
		// jika tidak ada datanya
		return nil, errors.New("record account with email not found")
	}
}

// method Get Account by username
func (a *AccountRepository) GetByUsername(ctx context.Context, username string) (*entity.Account, error) {
	query := "SELECT id, email, username, password, otp, expired_otp, created_at, user_id FROM accounts WHERE username=? LIMIT 1"
	row, err := a.DB.QueryContext(ctx, query, username)

	// jika ada error ketika proses query
	if err != nil {
		return nil, err
	}

	defer row.Close()

	account := entity.Account{}

	if row.Next() {
		// scan hasil query ke object account
		err := row.Scan(&account.ID, &account.Email, &account.Username, &account.Password, &account.OTP, &account.ExpiredOTP, &account.CreatedAt, &account.UserId)

		// jika ada error ketika proses scan
		if err != nil {
			return nil, err
		}

		// success scan -> return
		return &account, nil
	} else {
		// jika tidak ada datanya
		return nil, errors.New("record with username not found")
	}
}

// method Delete Account by ID
func (a *AccountRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM accounts WHERE id=?"
	_, err := a.DB.ExecContext(ctx, query, id)

	// jika ada error ketika delete
	if err != nil {
		return err
	}

	// success delete
	return nil
}

// method Update password by ID
func (a *AccountRepository) UpdatePassword(ctx context.Context, id int64, newPassword string) error {
	query := "UPDATE accounts SET password=? WHERE id=?"
	_, err := a.DB.ExecContext(ctx, query, newPassword, id)

	// jika ada kesalahan ketika update password
	if err != nil {
		return err
	}

	// success update password
	return nil
}

// method Update OTP by ID
func (a *AccountRepository) UpdateOTP(ctx context.Context, entity *entity.Account) error {
	query := "UPDATE accounts SET otp=?, expired_otp=? WHERE id=?"
	_, err := a.DB.ExecContext(ctx, query, entity.OTP, entity.ExpiredOTP, entity.ID)

	// jika ada kesalahan ketika update otp
	if err != nil {
		return err
	}

	// success update otp
	return nil
}

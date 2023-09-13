package repository

import (
	"cms/model/dto"
	"cms/model/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// create object UserRepository
type UserRepository struct {
	DB     *sql.DB
	Entity *entity.User
}

// create a function provider to create new object UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB:     db,
		Entity: &entity.User{},
	}
}

// method Insert
func (u *UserRepository) Insert(ctx context.Context, entity *entity.User) (*entity.User, error) {
	query := "INSERT INTO users (first_name, last_name, identity_id, gender, address) VALUES (?, ?, ?, ?, ?)"
	resultInsert, err := u.DB.ExecContext(ctx, query, entity.FirstName, entity.LastName, entity.IdentityId, entity.Gender, entity.Address)
	if err != nil {
		return nil, err
	}

	// get inserted ID
	id, err := resultInsert.LastInsertId()
	if err != nil {
		return nil, err
	}

	// set id to entity.ID
	entity.ID = id

	// success inserted -> return
	return entity, nil
}

// method Get by Id
func (u *UserRepository) GetById(ctx context.Context, id int64) (*entity.User, error) {
	query := "SELECT id, first_name, last_name, identity_id, gender, address, created_at FROM users WHERE id=? LIMIT 1"
	row, err := u.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create initiate object
	user := entity.User{}

	if row.Next() {
		// ada datanya
		// scan to object user
		row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.IdentityId, &user.Gender, &user.Address, &user.CreatedAt)

		// return
		return &user, nil
	} else {
		// jika tidak ada datanya
		return nil, errors.New("record not found")
	}
}

// method Delete by Id
func (u *UserRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id=?"
	_, err := u.DB.ExecContext(ctx, query, id)

	// jika ada kesalahan saat delete
	if err != nil {
		return err
	}

	// success delete
	return nil
}

// method Get Users by email
func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*dto.UsersByEmail, error) {
	query := "SELECT a.id, a.email, a.username, a.password, u.id, u.first_name, u.last_name, u.gender, u.address FROM accounts a INNER JOIN users u ON u.id = a.user_id WHERE a.email = ?"
	row, err := u.DB.QueryContext(ctx, query, email)

	// jika ada error saat proses query
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create object
	user := dto.UsersByEmail{}

	if row.Next() {
		// proses scan hasil query ke object
		err = row.Scan(&user.Account.Id, &user.Account.Email, &user.Account.Username, &user.Account.Password, &user.User.Id, &user.User.FirstName, &user.User.LastName, &user.User.Gender, &user.User.Address)

		// jika ada kesalah saat scan
		if err != nil {
			errString := fmt.Sprintf("error when scan : %v", err.Error())
			return nil, errors.New(errString)
		}
	} else {
		// jika tidak ada datanya
		return nil, errors.New("record not found")
	}

	// success get data -> return
	return &user, nil
}

// method Get All users
func (u *UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	query := "SELECT id, first_name, last_name, identity_id, gender, address, created_at FROM users"
	rows, err := u.DB.QueryContext(ctx, query)

	// jika ada error ketika query
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		// proses scan hasil query ke object
		var user entity.User
		rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.IdentityId, &user.Gender, &user.Address, &user.CreatedAt)

		users = append(users, user)
	}

	return users, nil
}

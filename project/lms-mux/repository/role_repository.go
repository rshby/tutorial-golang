package repository

import (
	"context"
	"errors"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB      *gorm.DB
	General *GeneralRepository[entity.Role]
}

func NewRoleRepository(db *gorm.DB, general *GeneralRepository[entity.Role]) *RoleRepository {
	return &RoleRepository{
		DB:      db,
		General: general,
	}
}

// method GetById role
func (r *RoleRepository) GetById(ctx context.Context, id int) (entity.Role, error) {
	role := entity.Role{}

	result := r.DB.WithContext(ctx).Model(&entity.Role{}).Preload("UserRoles").Where("roles.id=?", id).First(&role)
	if result.Error != nil {
		// jika data tidak ditemukan -> return with error not found
		return entity.Role{}, result.Error
	}

	// success get data role By Id
	// return
	return role, nil
}

// method GetAll roles
func (r *RoleRepository) GetAll(ctx context.Context) ([]entity.Role, error) {
	roles := []entity.Role{}

	result := r.DB.WithContext(ctx).Model(&entity.Role{}).Preload("UserRoles").Find(&roles)
	if result.RowsAffected == 0 {
		// jika data tidak ditemukan -> return with error not found
		return []entity.Role{}, errors.New("record roles not found")
	}

	// success get all data roles
	// return
	return roles, nil
}

// method get role by Name
func (r *RoleRepository) GetByName(ctx context.Context, name string) (entity.Role, error) {
	var role entity.Role

	result := r.DB.WithContext(ctx).Model(&entity.Role{}).Preload("UserRoles").Where("roles.name=?", name).First(&role)
	if result.Error != nil {
		// jika datanya tidak ada -> return with error data not found
		return entity.Role{}, result.Error
	}

	// success get data role by name
	return role, nil
}

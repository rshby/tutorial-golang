package repository

import (
	"context"
	"errors"
	"lms-mux/model/entity"

	"gorm.io/gorm"
)

type UserRoleRepository struct {
	DB      *gorm.DB
	General *GeneralRepository[entity.UserRole]
}

func NewuserRoleRepository(db *gorm.DB, general *GeneralRepository[entity.UserRole]) *UserRoleRepository {
	return &UserRoleRepository{
		DB:      db,
		General: general,
	}
}

// method get userRole by Id
func (u *UserRoleRepository) GetById(ctx context.Context, id int) (entity.UserRole, error) {
	var userRole entity.UserRole

	result := u.DB.WithContext(ctx).Model(&entity.UserRole{}).Preload("User").Preload("Role").Where("user_roles.id=?", id).First(&userRole)
	if result.Error != nil {
		// return with error data not found
		return entity.UserRole{}, result.Error
	}

	// success get data userRole by Id
	// return
	return userRole, nil
}

// method get data userRoles by user_id
func (u *UserRoleRepository) GetUserRolesByUserId(ctx context.Context, userId int) ([]entity.UserRole, error) {
	var userRoles []entity.UserRole

	result := u.DB.WithContext(ctx).Model(&entity.UserRole{}).Preload("User").Preload("Role").Where("user_roles.user_id=?", userId).Find(&userRoles)
	if result.RowsAffected == 0 {
		// return with error data not found
		return []entity.UserRole{}, errors.New("record user_roles not found")
	}

	// success get all userRoles by user_id
	// return
	return userRoles, nil
}

// method get data userrole by user_id and role_id
func (u *UserRoleRepository) GetByUserIdAndRoleId(ctx context.Context, userId int, roleId int) (entity.UserRole, error) {
	var userRole entity.UserRole

	result := u.DB.WithContext(ctx).Model(&entity.UserRole{}).Preload("User").Preload("Role").Where("user_id=? AND role_id=?", userId, roleId).First(&userRole)
	if result.Error != nil {
		// return with error data not found
		return entity.UserRole{}, result.Error
	}

	// success get data
	// return
	return userRole, nil
}

// method getall userroles
func (u *UserRoleRepository) GetAll(ctx context.Context) ([]entity.UserRole, error) {
	var userRoles []entity.UserRole

	result := u.DB.WithContext(ctx).Model(&entity.UserRole{}).Preload("User").Preload("Role").Find(&userRoles)
	if result.RowsAffected == 0 {
		// jika data tidak ada -> return with error not found
		return []entity.UserRole{}, errors.New("record user-roles not found")
	}

	// success get all data user-roles
	// return
	return userRoles, nil
}

// method get data user-roles by role_id
func (u *UserRoleRepository) GetByRoleId(ctx context.Context, roleId int) ([]entity.UserRole, error) {
	var userRoles []entity.UserRole

	result := u.DB.WithContext(ctx).Model(&entity.UserRole{}).Preload("User").Preload("Role").Where("user_roles.role_id=?", roleId).Find(&userRoles)
	if result.RowsAffected == 0 {
		// return with error data not found
		return []entity.UserRole{}, errors.New("record user-role not found")
	}

	// success get data user-role by role_id
	// return
	return userRoles, nil
}

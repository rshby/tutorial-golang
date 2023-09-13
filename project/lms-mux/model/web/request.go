package web

import "time"

type RequestProvinceInsert struct {
	Name string `json:"name" validate:"required"`
}

type RequestProvinceUpdate struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type RequestCityInsert struct {
	Name       string `json:"name" validate:"required"`
	ProvinceId int    `json:"province_id" validate:"required"`
}

type RequestCityUpdate struct {
	Id         int    `json:"id"`
	Name       string `json:"name" validate:"required"`
	ProvinceId int    `json:"province_id" validate:"required"`
}

type RequestDistrictInsert struct {
	Name   string `json:"name" validate:"required"`
	CityId int    `json:"city_id" validate:"required"`
}

type RequestDistrictUpdate struct {
	Id     int    `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	CityId int    `json:"city_id" validate:"required"`
}

type RequestSubDistrictInsert struct {
	Name       string `json:"name" validate:"required"`
	ZipCode    string `json:"zip_code" validate:"required"`
	DistrictId int    `json:"district_id" validate:"required"`
}

type RequestSubDistrictUpdate struct {
	Id         int    `json:"id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	ZipCode    string `json:"zip_code" validate:"required"`
	DistrictId int    `json:"district_id" validate:"required"`
}

type RequestAddressInsert struct {
	Street        string `json:"street" validate:"required"`
	SubDistrictId int    `json:"subdistrict_id" validate:"required"`
}

type RequestAddressUpdate struct {
	Id            int    `json:"id" validate:"required"`
	Street        string `json:"street" validate:"required"`
	SubDistrictId int    `json:"subdistrict_id" validate:"required"`
}

type RequestUniversityInsert struct {
	Name          string `json:"name" validate:"required"`
	Street        string `json:"street" validate:"required"`
	SubDistrictId int    `json:"subdistrict_id" validate:"required"`
}

type RequestuniversityUpdate struct {
	Id            int    `json:"id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Street        string `json:"street" validate:"required"`
	SubDistrictId int    `json:"subdistrict_id" validate:"required"`
}

// insert education
type RequestEducationInsert struct {
	Major        string `json:"major" validate:"required"`
	Level        string `json:"level" validate:"required"`
	UniversityId int    `json:"university_id" validate:"required"`
}

// update education
type RequestEducationUpdate struct {
	Id           int    `json:"id" validate:"required"`
	Major        string `json:"major" validate:"required"`
	Level        string `json:"level" validate:"required"`
	UniversityId int    `json:"university_id" validate:"required"`
}

// insert user
type RequestUserInsert struct {
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	Gender      string    `json:"gender" validate:"required,oneof=L P"`
	BirthDate   time.Time `json:"birth_date" validate:"required"`
	AddressId   int       `json:"address_id"`
	EducationId int       `json:"education_id"`
}

// update user
type RequestUserUpdate struct {
	Id          int       `json:"id" validate:"required"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	Gender      string    `json:"gender" validate:"required,oneof=L P"`
	BirthDate   time.Time `json:"birth_date" validate:"required"`
	AddressId   int       `json:"address_id"`
	EducationId int       `json:"education_id"`
}

// insert account
type RequestAccountInsert struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserId   int    `json:"user_id" validate:"required"`
}

// update account
type RequestUpdateAccount struct {
	Id       int    `json:"id" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserId   int    `json:"user_id" validate:"required"`
}

// insert role
type RequestRoleInsert struct {
	Name string `json:"name" validate:"required"`
}

// update role
type RequestRoleUpdate struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// insert user-role
type RequestUserRoleInsert struct {
	UserId int `json:"user_id" validate:"required"`
	RoleId int `json:"role_id" validate:"required"`
}

// update user-role
type RequestUserRoleUpdate struct {
	Id     int `json:"id" validate:"required"`
	UserId int `json:"user_id" validate:"required"`
	RoleId int `json:"role_id" validate:"required"`
}

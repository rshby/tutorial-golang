package entity

import "time"

// berisi entity

type Province struct {
	Id   int    `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Name string `gorm:"column:name" json:"name,omitempty"`

	// define relation
	Cities *[]City `gorm:"foreignKey:ProvinceId;references:Id" json:"cities,omitempty"`
}

type City struct {
	Id         int    `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	Name       string `gorm:"column:name" json:"name,omitempty"`
	ProvinceId int    `gorm:"column:province_id" json:"province_id,omitempty"`

	// define relation
	Province  *Province   `gorm:"foreignKey:ProvinceId" json:"province,omitempty"`
	Districts *[]District `gorm:"foreignKey:CityId;references:Id" json:"districts,omitempty"`
}

type District struct {
	Id     int    `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	Name   string `gorm:"column:name" json:"name,omitempty"`
	CityId int    `gorm:"column:city_id" json:"city_id,omitempty"`

	// define relation
	City         City           `gorm:"foreignKey:CityId" json:"city,omitempty"`
	SubDistricts *[]SubDistrict `gorm:"foreignKey:DistrictId;references:Id" json:"subdistricts,omitempty"`
}

type SubDistrict struct {
	Id         int    `gorm:"column:id;autoIncrement;primariKey" json:"id,omitempty"`
	Name       string `gorm:"column:name" json:"name,omitempty"`
	ZipCode    string `gorm:"zip_code" json:"zip_code,omitempty"`
	DistrictId int    `gorm:"column:district_id" json:"district_id,omitempty"`

	// define relation
	District  *District  `gorm:"foreignKey:DistrictId" json:"district,omitempty"`
	Addresses *[]Address `gorm:"foreignKey:SubDistrictId;references:Id" json:"addresses,omitempty"`
}

type Address struct {
	Id            int    `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	Street        string `gorm:"column:street" json:"street,omitempty"`
	SubDistrictId int    `gorm:"column:subdistrict_id" json:"subdistrict_id,omitempty"`

	// define relation
	SubDistrict  *SubDistrict  `gorm:"foreignKey:SubDistrictId" json:"subdistrict,omitempty"`
	User         *[]User       `gorm:"foreignKey:AddressId;references:Id" json:"users,omitempty"`
	Universities *[]University `gorm:"foreignKey:AddressId;references:Id" json:"universities,omitempty"`
}

type Education struct {
	Id           int    `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	Major        string `gorm:"column:major" json:"major,omitempty"`
	Level        string `gorm:"column:level" json:"level,omitempty"`
	UniversityId int    `gorm:"column:university_id" json:"university_id,omitempty"`

	// define relation
	Users      *[]User     `gorm:"foreignKey:EducationId;references:Id" json:"users,omitempty"`
	University *University `gorm:"foreignKey:UniversityId" json:"university,omitempty"`
}

type University struct {
	Id        int    `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	Name      string `gorm:"column:name" json:"name,omitempty"`
	AddressId int    `gorm:"column:address_id" json:"address_id,omitempty"`

	// define relation
	Educations *[]Education `gorm:"foreignKey:UniversityId;references:Id" json:"educations,omitempty"`
	Address    *Address     `gorm:"foreignKey:AddressId" json:"address,omitempty"`
}

type User struct {
	Id          int       `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	FirstName   string    `gorm:"column:first_name" json:"first_name,omitempty"`
	LastName    string    `gorm:"column:last_name" json:"last_name,omitempty"`
	Gender      string    `gorm:"column:gender" json:"gender,omitempty"`
	BirthDate   time.Time `gorm:"column:birthdate" json:"birth_date,omitempty"`
	AddressId   int       `gorm:"column:address_id" json:"address_id,omitempty"`
	EducationId int       `gorm:"column:education_id" json:"education_id,omitempty"`

	// define relation
	Address      *Address      `gorm:"foreignKey:AddressId" json:"address,omitempty"`
	Account      *Account      `gorm:"foreignKey:UserId;references:Id" json:"account,omitempty"`
	Education    *Education    `gorm:"foreignKey:EducationId" json:"education,omitempty"`
	TakenClasses *[]TakenClass `gorm:"foreignKey:UserId;references:Id" json:"takenclasses,omitempty"`
	UserRoles    *[]UserRole   `gorm:"foreignKey:UserId;references:Id" json:"user_roles,omitempty"`
}

type Account struct {
	Id       int    `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	Email    string `gorm:"column:email;unique" json:"email,omitempty"`
	Password string `gorm:"column:password" json:"password,omitempty"`
	UserId   int    `gorm:"column:user_id" json:"user_id,omitempty"`

	// define relation
	User *User `gorm:"foreignKey:UserId" json:"user,omitempty"`
}

type UserRole struct {
	Id     int `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	UserId int `gorm:"column:user_id" json:"user_id,omitempty"`
	RoleId int `gorm:"column:role_id" json:"role_id,omitempty"`

	// define relation
	User *User `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Role *Role `gorm:"foreignKey:RoleId" json:"role,omitempty"`
}

type Role struct {
	Id   int    `gorm:"column:id;autoIncrement;primaryKey" json:"id,omitempty"`
	Name string `gorm:"column:name" json:"name,omitempty"`

	// define relation
	UserRoles *[]UserRole `gorm:"foreignKey:RoleId;references:Id" json:"user_roles,omitempty"`
}

type ClassCategory struct {
	Id    int    `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Name  string `gorm:"name" json:"name,omitempty"`
	Level int    `gorm:"column:level" json:"level,omitempty"`

	// define relation
	Classes *[]Class `gorm:"foreignKey:ClassCategoryId;references:Id" json:"classes,omitempty"`
}

type Class struct {
	Id              int    `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Name            string `gorm:"column:name" json:"name,omitempty"`
	Description     string `gorm:"column:description" json:"description,omitempty"`
	TotalChapter    int    `gorm:"column:total_chapter" json:"total_chapter,omitempty"`
	ClassCategoryId int    `gorm:"column:classcategory_id" json:"classcategory_id,omitempty"`

	// define relation
	ClassCategory *ClassCategory `gorm:"foreignKey:ClassCategoryId" json:"classcategory,omitempty"`
	TakenClasses  *[]TakenClass  `gorm:"foreignKey:ClassId;references:Id" json:"takenclasses,omitempty"`
	Sections      *[]Section     `gorm:"foreignKey:ClassId;references:Id" json:"sections,omitempty"`
}

type TakenClass struct {
	Id              int  `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	ProgressChapter int  `gorm:"column:progress_chapter" json:"progress_chapter,omitempty"`
	IsDone          bool `gorm:"column:is_done" json:"is_done,omitempty"`
	UserId          int  `gorm:"column:user_id" json:"user_id,omitempty"`
	ClassId         int  `gorm:"column:class_id" json:"class_id,omitempty"`

	// define relation
	User  *User  `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Class *Class `gorm:"foreignKey:ClassId" json:"class,omitempty"`
}

type Section struct {
	Id      int    `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Title   string `gorm:"column:title" json:"title,omitempty"`
	Chapter int    `gorm:"column:chapter" json:"chapter,omitempty"`
	Content string `gorm:"column:content" json:"content,omitempty"`
	ClassId int    `gorm:"column:class_id" json:"class_id,omitempty"`

	// define relation
	Class *Class `gorm:"foreignKey:ClassId" json:"class,omitempty"`
}

package web

import "lms-mux/model/entity"

type ResponseJSON struct {
	StatusCode int    `json:"status_code,omitempty"`
	Status     string `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
}

type ResponseFullAddress struct {
	Id          int    `json:"id,omitempty"`
	Street      string `json:"street,omitempty"`
	SubDistrict string `json:"sub_district,omitempty"`
	ZipCode     string `json:"zip_code,omitempty"`
	District    string `json:"district,omitempty"`
	City        string `json:"city,omitempty"`
	Province    string `json:"province,omitempty"`
}

type ResponseUniversity struct {
	Id         int                 `json:"id,omitempty"`
	Name       string              `json:"name,omitempty"`
	Address    ResponseFullAddress `json:"address,omitempty"`
	Educations *[]entity.Education `json:"educations,omitempty"`
}

type ResponseEducation struct {
	Id         int    `json:"id,omitempty"`
	Major      string `json:"major,omitempty"`
	Level      string `json:"level,omitempty"`
	University string `json:"university,omitempty"`
}

type ResponseUser struct {
	Id        int                  `json:"id,omitempty"`
	FirstName string               `json:"first_name,omitempty"`
	LastName  string               `json:"last_name,omitempty"`
	Gender    string               `json:"gender,omitempty"`
	BirthDate string               `json:"birth_date,omitempty"`
	Account   *entity.Account      `json:"account,omitempty"`
	Address   *ResponseFullAddress `json:"address,omitempty"`
	Education *ResponseEducation   `json:"education,omitempty"`
}

type ResponseAccount struct {
	Id       int           `json:"id,omitempty"`
	Email    string        `json:"email,omitempty"`
	Password string        `json:"password,omitempty"`
	User     *ResponseUser `json:"user"`
}

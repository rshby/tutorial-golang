package router

import (
	"lms-mux/handler"
	"lms-mux/helper"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(
	provinceHandler *handler.ProvinceHandler,
	cityHandler *handler.CityHandler,
	districtHandler *handler.DistrictHandler,
	subdistrictHandler *handler.SubDistrictHandler,
	addressHandler *handler.AddressHandler,
	univHandler *handler.UniversityHandler,
	educationHandler *handler.EducationHandler,
	userHandler *handler.UserHandler,
	accountHandler *handler.AccountHandler,
	roleHandler *handler.RoleHandler,
	userRoleHandler *handler.UserRoleHandler) *mux.Router {
	r := mux.NewRouter()

	r.MethodNotAllowedHandler = helper.NotAllowedMethod()
	r.NotFoundHandler = helper.NotFoundMethod()

	// buat endpoint
	// province
	r.HandleFunc("/api/province", provinceHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/province/{id}", provinceHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/province/{id}", provinceHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/provinces", provinceHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/province/{id}", provinceHandler.GetById).Methods(http.MethodGet)

	// cities
	r.HandleFunc("/api/city", cityHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/city", cityHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/city/{id}", cityHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/cities", cityHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/city/{id}", cityHandler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/province/{province_id}/cities", cityHandler.GetAllCitiesByProvinceId).Methods(http.MethodGet)

	// distrits
	r.HandleFunc("/api/district", districtHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/district", districtHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/district/{id}", districtHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/districts", districtHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/district/{id}", districtHandler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/city/{id}/districts", districtHandler.GetAllDistrictsByCityId).Methods(http.MethodGet)

	// sub_districts
	r.HandleFunc("/api/subdistrict", subdistrictHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/subdistrict", subdistrictHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/subdistrict/{id}", subdistrictHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/subdistricts", subdistrictHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/subdistrict/{id}", subdistrictHandler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/district/{id}/subdistricts", subdistrictHandler.GetByDistrictId).Methods(http.MethodGet)

	// address
	r.HandleFunc("/api/address", addressHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/address", addressHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/address/{id}", addressHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/addresses", addressHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/address/{id}", addressHandler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/full-address/{id}", addressHandler.GetFullAddressById).Methods(http.MethodGet)

	// university
	r.HandleFunc("/api/university", univHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/university", univHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/university/{id}", univHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/universities", univHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/university/{id}", univHandler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/university/{id}/educations", univHandler.GetAllEducationsByUniversityId).Methods(http.MethodGet)

	// education
	r.HandleFunc("/api/education", educationHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/education", educationHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/education/{id}", educationHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/educations", educationHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/education/{id}", educationHandler.GetById).Methods(http.MethodGet)

	// user
	r.HandleFunc("/api/user", userHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/user/insert-mass", userHandler.CreateMassUser).Methods(http.MethodPost)
	r.HandleFunc("/api/user", userHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/user/{id}", userHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/users", userHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/user/{id}", userHandler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/user-polos/{id}", userHandler.GetByIdFromEntity).Methods(http.MethodGet)

	// account
	r.HandleFunc("/api/account", accountHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/account", accountHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/account/{id}", accountHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/accounts", accountHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/account/{id}", accountHandler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/account", accountHandler.GetByEmail).Methods(http.MethodGet)

	// role
	r.HandleFunc("/api/role", roleHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/role", roleHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/role/{id}", roleHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/roles", roleHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/role/{id}", roleHandler.GetById).Methods(http.MethodGet)

	// user-role
	r.HandleFunc("/api/user-role", userRoleHandler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/user-role", userRoleHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/user-role/{id}", userRoleHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/user-roles", userRoleHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/user-role/{id}", userRoleHandler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/user-role", userRoleHandler.GetByUserIdAndRoleId).Methods(http.MethodGet)

	// class-category

	// class

	// section

	// takenclass
	// -> daftar class
	// -> next chapter

	// login & logout
	// -> login
	// -> logout
	// -> forgot password

	return r
}

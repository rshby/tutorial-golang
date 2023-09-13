package testing

import (
	"lms-mux/handler"
	"lms-mux/model/entity"
	"lms-mux/repository"
	"lms-mux/router"
	"lms-mux/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var handler_test = NewRouterTest()

func NewDatabaseTesting() *gorm.DB {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/lms_test?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	db.AutoMigrate(&entity.Account{}, &entity.Address{}, &entity.City{}, &entity.Class{}, &entity.ClassCategory{}, &entity.District{}, &entity.Education{}, &entity.Province{}, &entity.Role{}, &entity.Section{}, &entity.SubDistrict{}, &entity.TakenClass{}, &entity.University{}, &entity.User{}, &entity.UserRole{})

	return db
}

func NewRouterTest() *mux.Router {
	db := NewDatabaseTesting()
	validate := validator.New()

	provinceRepository := repository.NewProvinceRepository(db)
	provinceService := service.NewProvinceService(provinceRepository, validate)
	provinceHandler := handler.NewProvinceHandler(provinceService)

	cityRepository := repository.NewCityRepository(db)
	cityService := service.NewCityService(cityRepository, provinceRepository, validate)
	cityHandler := handler.NewCityHandler(cityService)

	districtRepository := repository.NewDistrictRepository(db)
	districtService := service.NewDistrictService(districtRepository, cityRepository, validate)
	districtHandler := handler.NewDistrictHandler(districtService)

	subDistrict := entity.NewEntity[entity.SubDistrict]()
	subDistrictGeneralRepository := repository.NewGeneralRepository(db, subDistrict)
	subDistrictRepository := repository.NewSubDistrictRepository(db)
	subDistrictService := service.NewSubDistrictService(subDistrictGeneralRepository, subDistrictRepository, districtRepository, validate)
	subDistrictHandler := handler.NewSubDistrictHandler(subDistrictService)

	// address
	address := entity.NewEntity[entity.Address]()
	addressGeneralRepository := repository.NewGeneralRepository(db, address)
	addressRepository := repository.NewAddressRepository(db, addressGeneralRepository)
	addressService := service.NewAddressService(addressGeneralRepository, addressRepository, validate)
	addressHandler := handler.NewAddressHandler(addressService)

	// university
	univRepo := repository.NewUniversityRepository(db)
	univGeneralRepo := repository.NewGeneralRepository(db, &entity.University{})
	univService := service.NewUniversityService(univGeneralRepo, univRepo, addressGeneralRepository, validate)
	univHandler := handler.NewUniversityHandler(univService)

	// education
	educationEntity := entity.NewEntity[entity.Education]()
	educationGeneralRepo := repository.NewGeneralRepository(db, educationEntity)
	educationRepo := repository.NewEducationRepository(db, educationGeneralRepo)
	educationService := service.NewEducationService(educationRepo, validate)
	educationHandler := handler.NewEducationHandler(educationService)

	// account
	account := entity.NewEntity[entity.Account]()
	accountGeneralRepo := repository.NewGeneralRepository(db, account)
	accountRepo := repository.NewAccountRepository(db, accountGeneralRepo)
	accountService := service.NewAccountService(accountRepo, validate)
	accountHandler := handler.NewAccountHandler(accountService)

	// takenclass
	takenclass := entity.NewEntity[entity.TakenClass]()
	takenClassGeneralRepo := repository.NewGeneralRepository(db, takenclass)
	takenClassRepo := repository.NewTakenClassRepository(db, takenClassGeneralRepo)

	// userRole
	userRole := entity.NewEntity[entity.UserRole]()
	userRoleGeneralRepo := repository.NewGeneralRepository(db, userRole)
	userRoleRepo := repository.NewuserRoleRepository(db, userRoleGeneralRepo)
	userRoleService := service.NewUserRoleService(userRoleRepo, validate)
	userRoleHandler := handler.NewUserRoleHandler(userRoleService)

	// user
	user := entity.NewEntity[entity.User]()
	userGeneralRepo := repository.NewGeneralRepository(db, user)
	userRepo := repository.NewUserRepository(db, userGeneralRepo)
	userService := service.NewUserService(userRepo, accountRepo, takenClassRepo, userRoleRepo, validate)
	userHandler := handler.NewUserHandler(userService)

	// role
	role := entity.NewEntity[entity.Role]()
	roleGeneralRepo := repository.NewGeneralRepository(db, role)
	roleRepo := repository.NewRoleRepository(db, roleGeneralRepo)
	roleService := service.NewRoleService(roleRepo, validate)
	roleHandler := handler.NewRoleHandler(roleService)

	// router
	muxRouter := router.NewRouter(provinceHandler, cityHandler, districtHandler, subDistrictHandler, addressHandler, univHandler, educationHandler, userHandler, accountHandler, roleHandler, userRoleHandler)

	return muxRouter
}

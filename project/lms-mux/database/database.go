package database

import (
	"lms-mux/model/entity"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/lms?parseTime=true"), &gorm.Config{})
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

package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
	"training/app/model/entity"
)

func NewConnection() *gorm.DB {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("root:@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", db_host, db_port, db_name)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf(fmt.Sprintf("error connection to database: %v", err.Error()))
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(50)
	sqlDb.SetConnMaxIdleTime(10 * time.Minute)
	sqlDb.SetConnMaxLifetime(1 * time.Hour)

	db.AutoMigrate(&entity.User{}, &entity.Book{})

	return db
}

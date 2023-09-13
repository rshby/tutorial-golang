package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// function koneksi ke database
func ConnectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_golang?parseTime=true")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Sukses Koneksi Database!")
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(100 * time.Second)
	db.SetConnMaxLifetime(100 * time.Minute)

	return db
}

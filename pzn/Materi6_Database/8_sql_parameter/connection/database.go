package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang?parseTime=true")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Koneksi Database Sukses!")
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(100 * time.Minute)
	db.SetConnMaxIdleTime(100 * time.Second)

	return db
}

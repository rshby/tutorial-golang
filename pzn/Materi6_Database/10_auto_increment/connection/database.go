package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/belajar_golang?parseTime=true")

	if err != nil {
		panic("Koneksi Gagal " + err.Error())
	} else {
		fmt.Println("Koneksi Sukses")
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(100 * time.Second)
	db.SetConnMaxLifetime(100 * time.Minute)

	return db
}

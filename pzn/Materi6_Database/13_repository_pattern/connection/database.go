package connection

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// function untuk koneksi ke database
func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/belajar_golang?parseTime=true")
	if err != nil {
		panic("Gagal Koneksi ke Database = " + err.Error())
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(100 * time.Second)
	db.SetConnMaxLifetime(100 * time.Minute)

	return db
}

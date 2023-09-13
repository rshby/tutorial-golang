package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
=== Membuat koneksi ke database ===
-> hal pertama yang akan kita lakukan ketika membuat aplikasi yang akan menggunakan database adalah melakukan koneksi ke databasenya
-> untuk melakukan koneksi ke database di golang, kita bisa membuat object sql.DB menggunakan function sql.Open(driver, dataSourceName)
-> untuk menggunakan database MySql, kita menggunakan driver mysql
-> sedangkan untuk dataSourceName, tiap database biasanya punya cara penulisan masing-masing misal di MySQL kita bisa menggunakan username:password@tcp(host:port)/db_name
-> jika object sql.DB sudah tidak digunakan lagi, disarankan untuk menutupnya menggunakan function close()
*/

// buat function untuk Koneksi Database
func ConnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/toko_online")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Koneksi database sukses!")
	}
	// gunakan db
	db.Close()
}

func main() {
	ConnectDB()
}

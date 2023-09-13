package main

import (
	"4_database_pooling/database"
	_ "github.com/go-sql-driver/mysql"
)

/*
=== Database Pooling ===
-> sql.DB di golang sebenarnya bukanlah sebuah koneksi ke database
-> melainkan sebuah pool ke database, atau dikenal dengan konsep Database Pooling
-> di dalam sql.DB golang melakukan management koneksi ke database secara otomatis. Hal ini menjadikan kita tidak perlu melakukan management koneksi database secara manual
-> dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal koneksi yang digunakan oleh golang. sehingga kita tidak membanjiri koneksi database, karena biasanya ada batas maksimal koneksi yang bisa ditangani oleh database yang kita gunakan
*/

func main() {
	database.GetConnection()
}

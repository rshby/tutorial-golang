package database

import "fmt"

/**
ketika ingin membuat function yang otomatis berjalan ketika sebuah package dipanggil, maka dapat dibuat function dengan nama ini
seperti constructor dalam class OOP
*/

// membuat variabel
var connection string

// membuat function init -> akan dijalankan saat package diakses
func init() {
	fmt.Println("function init dijalankan")
	connection = "Mysql"
}

// membuat function return connection
func GetDatabase() string {
	return connection
}

package main

import "10_auto_increment/response"

/*
AUTO INCREMENT
-> kadang kita membuat sebuah tabel dengan id auto increment
-> dan kadang pula, kita ingin mengambil data id yang sudah kita insert ke dalam mysql
-> sebenarnya kita bisa melakukan query ulang ke database menggunakan SELECT LAST_INSERT_ID()
-> tapi untungnya di golang ada cara yang lebih mudah
-> kita bisa menggunakan function (Result)LastInsertId() untuk mendapatkan Id terakhir yang dibuat secara auto increment
-> Result adalah object yang dikembalikan ketika kita menggunakan function Exec()
*/

func main() {
	// insert ke database
	response.InsertComment("reoshby@gmail.com", "komen2")
}

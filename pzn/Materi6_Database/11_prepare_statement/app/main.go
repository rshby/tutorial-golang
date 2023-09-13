package main

import "11_prepare_statement/response"

/*
QUERY atau Exec dengan Parameter
-> saat kita menggunakan function Query atau Exec yang menggunakan parameter, sebenarnya implementasi dibawahnya Prepare Statement
-> jadi tahapan pertama statementnya disiapkan terlebih dahulu, setelah itu baru diisi dengan parameter
-> kadang ada kasus kita ingin melakukan beberapa hal yang sama sekaligus, hanya berbeda parameternya. Misal insert data langsung banyak
-> pembuatan prepare statement bisa dilakukan dengan manual, tanpa harus menggunakan Query atau Exec dengan parameter
*/

/*
prepare Statement
-> saat kita membuat prepare statement, secara otomatis, akan mengenali koneksi database yang digunakan
-> sehingga ketika kita mengeksekusi prepare statement berkali-kali, maka akan menggunakan koneksi yang sama dan lebih efisien karena pembuatan prepare statementnya hanya sekali diawal saja
-> jika menggunakan Query atau Exec dengan parameter, kita tidak bisa menjamin bahwa koneksi yang digunakan akan sama, oleh karena itu bisa jadi prepare statement akan selalu dibuat berkali-kali walaupun kita mengunakan SQL yang sama
-> untuk membuat prepare statement, kita menggunakan function (db).Prepare(ctx, sql)
-> prepare statement direpresentasikan dalam struct database/sql.Stmt
-> sama seperti resource lainnya, Stmt harus diClose() jika sudah tidak digunakan lagi
*/

func main() {
	// insert comment
	response.InsertComment()

}

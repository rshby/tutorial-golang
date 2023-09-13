package main

import "9_sql_dengan_parameter/response"

/*
SQL Dengan Parameter
-> sekarang kita sudah tau bahayanya sql injection
-> jika ada kebutuhan seperti itu, sebenarnya function Exec dan Query memiliki parameter tambahan yang bisa kita gunakan untuk memsubsitusi parameter dari sebuah function tersebut ke sql query yang kita buat
-> untuk menandai sebuah sql membutuhkan parameter, kita bisa gunakan karakter ? (tanda tanya)
*/

/*
CONTOH QUERY
-> SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
-> INSERT INTO user (username, password) VALUES (?, ?)
-> dan lain-lain...
*/

func main() {
	response.InsertData("reoshby", "12345678")
	response.GetUserByParams()
}

package main

import "8_sql_parameter/response"

/*
SQL dengan parameter
-> saat kita membuat aplikasi, kita tidak mungkin akan melakukan harcode perintah sql di kode golang kita
-> biasanya kita akan menerima input dari data user, lalu membuat perintah sql dari input user, dan mengirimnya menggunakan perintah sql
*/

/*
SQL Injection
-> adalah sebuah teknik yang menyalahgunakan sebuah celah keamanan yang terjadi dalam lapisan database sebuah aplikasi
-> biasanya, sql injenction dilakukan dengan mengirim input dari user dengan perintah yang salah, sehingga menyebabkan hasil SQL yang kita buat menjadi tidak valid
-> SQL injection ini sangat berbahaya, jika sampai kita salah membuat SQL, bisa jadi data kita tidak aman
*/

/*
SOLUSINYA
-> jangan membuat sql dengan secara manual menggabungkan string parameter
-> jika membutuhkan parameter ketika membuat sql, kita bisa menggunakan function Execute atau Query dengan parameter
*/

func main() {
	response.GetUserByParams()
}

package main

import "12_database_transaction/response"

/*
DATABASE TRANSACTION
-> salah satu fitur andalan di database adalah transaction
-> materi database transaction sudah dibahas dalam mysql
-> di course ini kita akan fokus menggunakan mysql transaction di golang
*/

/*
TRANSACTION DI GOLANG
-> secara default, semua perintah SQL yang kita kirim menggunakan Golang akan otomatis di commit, atau istilahnya auto commit
-> namun kita bisa menggunakan transaction sehingga SQL yang kita kirim tidak secara otomatis dicommit ke database
-> Untuk memulai transaksi, kita bisa menggunakan function (db).Begin(), dimana akan menghasilkan struct Tx yang merupakan representasi Transaction
-> Struct Tx ini yang akan kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua function di DB ada di Tx, seperti Exec, Query, atau Prepare
-> setelah selesai proses transaksi, kita bisa gunakan (Tx).Commit() untuk melakukan commit atau Rollback()
*/

func main() {
	response.InsertCommentData()
}

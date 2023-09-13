package main

import "5_eksekusi_query/controllers"

/*
=== Eksekusi Perintah SQL ===
-> saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan database menggunakan perintah SQL
-> di golang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan function db.ExecContext(context, sql, params)
-> ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti materi sebelumnya, dengan context kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengiriman query SQLnya
*/

func main() {
	controllers.InsertData()
}

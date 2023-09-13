package main

import "6_query_sql/controllers"

/*
=== Query SQL ===
-> untuk operasi SQL yang tidak membutuhkan hasil dapat menggunakan db.Exec, namun juga query yang dapat menghasilkan data kita dapat menggunakan db.QueryContext(context, sql, params)
*/

/*
=== Rows ===
-> hasil query function adalah sebuah struct sql.Rows
-> Rows digunakan untuk melakukan iterasi terhadap hasil dari query
-> kita bisa menggunakan function rows.Next() untuk melakukan iterasi terhadap data hasil query, jika return data false, artinya sudah tidak ada lagi data dalam rows
-> untuk membaca tiap data, kita bisa menggunakan rows.Scan(columns...)
-> jangan lupa setelah menggunakan rows ditutup dengan function rows.Close()
*/

func main() {
	controllers.GetCustomers()
}

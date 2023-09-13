package response

import (
	"12_database_transaction/connection"
	"context"
	"fmt"
	"strconv"
)

func InsertCommentData() {
	db := connection.ConnectDB12()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic("Gagal membuat transaction Tx = " + err.Error())
	}

	// do transaction
	sqlQuery := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	for i := 0; i < 10; i++ {
		email := "reoshby1@gmail.com"
		comment := "ini komen ke-" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, sqlQuery, email, comment)
		if err != nil {
			panic("Gagal Insert Menggunakan Transaction = " + err.Error())
		}

		// ambil id
		if lastId, err := result.LastInsertId(); err != nil {
			panic("Gagal Menampilkan Last Id = " + err.Error())
		} else {
			fmt.Println("Sukses Data Berhasil DiTambahkan dengan Id =", lastId)
		}
	}

	if err := tx.Rollback(); err != nil {
		panic("Error Commit Transaction = " + err.Error())
	}
}

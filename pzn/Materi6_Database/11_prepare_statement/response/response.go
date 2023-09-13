package response

import (
	"11_prepare_statement/connection"
	"context"
	"fmt"
	"strconv"
)

// function untuk Insert ke Database tabel Comment
func InsertComment() {
	db := connection.ConectDB11()
	defer db.Close()

	ctx := context.Background()

	// buat statement
	stmt, err := db.PrepareContext(ctx, "INSERT INTO comments (email, comment) VALUES (?, ?)")
	if err != nil {
		panic("GAGAL INSERT " + err.Error())
	} else {
		fmt.Println("SUKSES Prepare KE comments")
	}

	defer stmt.Close()

	// iterasi data yang akan diinsert
	for i := 3; i <= 10; i++ {
		email := "reoshby@gmail.com"
		komen := "ini komen ke " + strconv.Itoa(i)

		// proses eksekusi statement (berisi query insert ke database)
		result, err := stmt.ExecContext(ctx, email, komen)
		if err != nil {
			panic("GAGAL INSERT " + err.Error())
		}

		// ambil id terakhirny
		id, err := result.LastInsertId()
		if err != nil {
			panic("GAGAL GET LAST ID = " + err.Error())
		}

		fmt.Println("SUKSES INSERT ID TERAKHIR", id)
	}
}

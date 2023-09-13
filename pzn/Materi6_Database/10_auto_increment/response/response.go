package response

import (
	"10_auto_increment/connection"
	"context"
	"fmt"
)

// function yang digunakan untuk insert data
func InsertComment(inputEmail string, inputComment string) {
	ctx := context.Background()

	db := connection.ConnectDB()
	defer db.Close()

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, inputEmail, inputComment)
	if err != nil {
		panic("Gagal Insert ke Database : " + err.Error())
	} else {
		insertid, err := result.LastInsertId()
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Sukses insert id ", insertid)
	}
}

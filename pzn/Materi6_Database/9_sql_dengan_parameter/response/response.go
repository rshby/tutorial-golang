package response

import (
	"9_sql_dengan_parameter/connection"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// function untuk GET data by username and password
func GetUserByParams() {
	db := connection.ConnectToDB()
	defer db.Close()

	ctx := context.Background()
	username := "admin"
	password := "admin"

	query := "SELECT username, password FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println("USERNAME =", username, " PASSWORD =", password)
		}
	}
}

// function untuk INSERT data ke database
func InsertData(username string, password string) {
	db := connection.ConnectToDB()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO user (username, password) VALUES (?, ?)"

	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic("GAGAL INSERT = " + err.Error())
	} else {
		fmt.Println("SUKSES INSERT INTO DATABASE user")
	}
}

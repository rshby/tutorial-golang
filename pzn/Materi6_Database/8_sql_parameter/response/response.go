package response

import (
	"8_sql_parameter/connection"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// function untuk GET user by username and password
func GetUserByParams() {
	db := connection.ConnectToDB()
	defer db.Close()

	ctx := context.Background()
	username := "admin' ; #"
	password := "salah"
	query := "SELECT username, password FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println("USERNAME =", username, " Password =", password)
		}
	}

}

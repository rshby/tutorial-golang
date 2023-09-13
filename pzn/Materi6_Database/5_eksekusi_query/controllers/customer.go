package controllers

import (
	"5_eksekusi_query/database"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// function yang digunakan untuk insert data ke database
func InsertData() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES (3, 'Reo Hidayanti');")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sukses Insert ke Database")
	}
}

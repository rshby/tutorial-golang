package controllers

import (
	"6_query_sql/database"
	"context"
	"fmt"
)

// buat function untuk get data
func GetCustomers() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT * FROM customer;")
	if err != nil {
		panic(err)
	} else {
		// iterasi tiap data
		for rows.Next() {
			var id int
			var name string

			// scan datanya tiap iterasi
			err := rows.Scan(&id, &name)

			// apabila error
			if err != nil {
				panic(err)
			} else {
				// tampilkan data
				fmt.Println("Id =", id, "Nama =", name)
			}
		}
		defer rows.Close()
	}
}

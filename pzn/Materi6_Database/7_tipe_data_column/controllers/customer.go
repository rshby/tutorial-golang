package controllers

import (
	"7_tipe_data_column/database"
	"context"
	"database/sql"
	"fmt"
	"time"
)

// buat function untuk get data
func GetCustomers() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	// proses query
	rows, err := db.QueryContext(ctx, "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer;")

	if err != nil {
		panic(err)
	} else {
		// iterasi tiap data
		for rows.Next() {
			var id, balance int
			var name, email sql.NullString
			var rating float64
			var married bool
			var birth_date, created_at time.Time

			// scan datanya tiap iterasi
			err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)

			// apabila error
			if err != nil {
				panic(err)
			} else {
				// tampilkan data
				fmt.Println("Id =", id)
				fmt.Println("Nama =", name)
				fmt.Println("Email =", email)
				fmt.Println("Balance =", balance)
				fmt.Println("Rating =", rating)
				fmt.Println("Birth Date =", birth_date)
				fmt.Println("Married =", married)
				fmt.Println("Created At =", created_at)
				fmt.Println("")
			}
		}
		defer rows.Close()
	}
}

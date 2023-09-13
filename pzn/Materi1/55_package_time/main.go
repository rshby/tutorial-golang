package main

import (
	"fmt"
	"time"
)

/*
berisikan fungsionalitas manajemen waktu di golang
*/

func main() {
	// time.Now() -> mendapatkan waktu saat ini
	fmt.Println("waktu saat ini adalah", time.Now().Local())

	// mendapatkan hari -> time.Now().Local().Weekday()
	fmt.Println("hari saat ini adalah =", time.Now().Local().Weekday())

	// mendapatkan tanggal -> time.Now().Local().Day()
	fmt.Println("tanggal saat ini adalah =", time.Now().Local().Day())

	// mendapatkan bulan -> time.Now().Local().Month()
	fmt.Println("bulan saat ini adalah =", time.Now().Local().Month())

	// mendapatkan tahun -> time.Now().Local().Year()
	fmt.Println("tahun saat ini adalah =", time.Now().Local().Year())

	// mendapatkan jam -> time.Now().Local().Hour()
	fmt.Println("jam saat ini adalah =", time.Now().Local().Hour())

	// mendapatkan menit -> time.Now().Local().Minute()
	fmt.Println("menit saat ini adalah =", time.Now().Local().Minute())

	// time.Date(...) -> untu membuat waktu
	utc := time.Date(2022, 8, 27, 8, 30, 0, 0, time.Local)
	fmt.Println("membuat waktu =", utc)

	// time.Parse(layout, string) -> untuk memparsing waktu dari string
	layout := "2006-01-02"
	if parse, err := time.Parse(layout, "2022-08-27"); err != nil {
		fmt.Print("error =", err.Error())
	} else {
		fmt.Println("waktu =", parse)
	}
}

package main

import (
	"fmt"
)

// switch di Golang
func main() {
	// == 1. switch case normal ==
	fmt.Println("- 1. switch case normal -")

	nilai := 80
	switch nilai {
	case 100:
		fmt.Println("selamat mendapat A")
	case 80:
		fmt.Println("gapapa masih dapat B")
	}

	// == 2. case bisa dua kondisi ==
	fmt.Println("\n- 2. case bisa dua kondisi -")
	nilai = 60
	switch nilai {
	case 100:
		fmt.Println("selamat anda mendapatkan A")
	case 60, 50:
		fmt.Println("belajar lagi ya")
	}

	// == 3. menggunakan variabel di dalam scope switch
	fmt.Println("\n- 3. menggunakan variabel di dalam scope switch -")
	switch x := 10; x {
	case 10:
		fmt.Println("variabel x bernilai 10")
	case 5:
		fmt.Println("variabel x bernilai 5")
	}

	// == 4. default ==
	// default adalah case terakhir yang akan dijalankan jika semua case tidak dijalankan
	fmt.Println("\n- 4. default -")
	nilai = 99
	switch nilai {
	case 100:
		fmt.Println("selamat anda mendapatkan A")
	case 90:
		fmt.Println("selamat anda masih mendapatkan AB")
	default:
		fmt.Println("nilai anda", nilai, "tidak masuk dalam case")
	}

	// == 5. switch case tanpa statement ==
	fmt.Println("\n- 5. switch case tanpa statement -")
	nilai = 80
	switch {
	case nilai == 100:
		fmt.Println("nilai anda 100, anda mendapatkan A")
	case nilai == 80:
		fmt.Println("nilai anda 80, anda mendapatkan B")
	default:
		fmt.Println("nilai anda", nilai, "tidak masuk dalam case")
	}

	// == 6. switch case short statement ==
	fmt.Println("\n- 6. switch case short statement -")
	name := "muhammad reo sahobby"
	switch panjang := len(name); panjang > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama pendek")
	}

	// == 7. type case
	fmt.Println("\n- 7. type case -")
	var newData any
	newData = string("reo")
	switch newData.(type) {
	case bool:
		fmt.Println("variabel bertipe boolean")
	case string:
		fmt.Println("variabel bertipe string")
	case int:
		fmt.Println("variabel bertipe int")
	default:
		fmt.Println("error tidak ditemukan tipe data yang cocok")
	}
}

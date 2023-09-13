package main

import "fmt"

// if else di Golang
func main() {
	// == 1. if else ==
	fmt.Println("- 1. if else -")

	myName := "reo"
	if myName == "reo" {
		fmt.Println("ya benar, myName bernilai 'reo'")
	} else {
		fmt.Println("salah, nama bukan reo, tapi", myName)
	}

	// == 2. if else if ==
	fmt.Println("\n- 2. if else if -")

	nilaiUjian := 100
	if nilaiUjian >= 80 {
		fmt.Println("selamat anda mendapatkan A")
	} else if nilaiUjian >= 70 && nilaiUjian < 80 {
		fmt.Println("gapapa, masih dapet B")
	} else if nilaiUjian >= 60 {
		fmt.Println("gapapa, jelek dikit ga ngaruh masih dapet C")
	} else {
		fmt.Println("belajar lagi ya")
	}

	// == 3. if else menggunakan short statement ==
	fmt.Println("\n- 3. if else menggunakan short statement -")

	myName = "muhammad reo sahobby eka saputra"
	if panjangNama := len(myName); panjangNama >= 10 {
		fmt.Println("nama terlalu panjang, sebanyak", panjangNama, "karakter")
	} else {
		fmt.Println("nama tidak terlalu panjang, hanya", panjangNama, "karakter")
	}
}

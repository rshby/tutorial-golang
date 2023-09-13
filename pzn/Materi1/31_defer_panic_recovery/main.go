package main

import "fmt"

// defer funtion -> function yang dipanggil setelah function sebelumnya dijalankan (tidak peduli error/success)
// panic -> syntax yang akan menghentikan program ketika ketemu error (codingan setelah ketemu panic tidak akan dijalankan)
// recover -> syntax yg akan menangkap panic dan melanjutkan program (supaya tidak berhenti di error)

// defer function
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Aplikasi Error, Message =", message)
	}
	fmt.Println("Aplikasi Selesai...")
}

// function run aplikasi
func runApp(inputError bool) {
	defer endApp()
	fmt.Println("Aplikasi Dimulai")
	if inputError {
		panic("***Error***")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true) // jika error
	fmt.Println("hello world")

	runApp(false) // jika tidak error
	fmt.Println("Reo Sahobby")
}

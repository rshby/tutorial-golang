package main

import "fmt"

// interface kosong -> tidak memiliki kontrak : tipe data yang bisa menerima apapun

// contoh pada function dengan return interface kosong
func Get(inputNumber int) interface{} {
	if inputNumber == 1 {
		return 1
	} else if inputNumber == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	// panggil function
	data := Get(2) // tipa datanya interface{}
	fmt.Println(data)
}

package main

import (
	"fmt"
)

// type assertion -> kemampuan merubah tipe data dari any menjadi tipe data yang diinginkan
// sering kali digunakan untuk merubah tipe data any

// membuat function
func random() interface{} {
	return "OK"
}

func main() {
	// membuat variabel
	result := random()

	// merubah variabel dari yg tipe datanya interface{} ke string
	resultString := result.(string)

	fmt.Println("Isi Variabel resultString =", resultString)

	// Switch Case tipe data result
	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is String")
	case int:
		fmt.Println("value", value, "is int")
	default:
		fmt.Println("Unknown Data Type")
	}
}

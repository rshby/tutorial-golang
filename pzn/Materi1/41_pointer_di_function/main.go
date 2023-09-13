package main

import (
	"fmt"
)

/**
saat membuat function, parameter yang dikirim bersifat pass by value
untuk merubah data aslinya perlu menggunakan parameter pointer
untuk menggunakan pointer di function menggunakan *(bintang) di parameternya
*/

// membuat struct
type Address struct {
	ID      int
	City    string
	Provice string
	County  string
}

// buat variabel
var tes12 string

// membuat function untuk menambahkan/merubah country menjadi Indonesia -> input parameter tipe datanya ditambahkan *(bintang)
func SetCountryToIndonesia(inputobject *Address) {
	tes12 = "Haloo" // isi variabel akan tetap diganti karena bukan parameter inputan
	inputobject.County = "Indonesia"
}

func main() {
	// buat object dari struct Address
	address1 := Address{
		ID:      1,
		City:    "Jakarta Selatan",
		Provice: "DKI Jakarta",
	}

	fmt.Println("address1 adalah =", address1) // country belum ada

	fmt.Println("isi variabel tes12 =", tes12) // kosong

	// panggil function untuk menambahkan/merubah country -> input parameter menggunakan &(dan)
	SetCountryToIndonesia(&address1)

	fmt.Println("address1 setelah ditambahkan adalah =", address1) // country sudah diisi

	fmt.Println("isi variabel tes12 =", tes12) // sudah diisi pada function
}

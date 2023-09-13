package main

import "fmt"

/*
data struct yang diakses di dalam method sebenarnya bersifat pass by value
sangat direkomendasikan menggunakan pointer supaya pass by reference biar lebih hemat memory
*/

// buat struct Person
type Person struct {
	ID     int
	Nama   string
	Email  string
	Alamat string
}

// buat function Married() yg seolah-olah milih object -> menggunakan *(bintang) pada input structnya
func (inputObject *Person) Married() {
	inputObject.Nama = "Mr. " + inputObject.Nama
}

func main() {
	// buat object dari struct Person
	reo := Person{
		ID:     1,
		Nama:   "Reo Sahobby",
		Email:  "reoshby@gmail.com",
		Alamat: "Jl.Swadaya 1 No.57, Rt05/Rw05, Ragunan, Pasar Minggu, Jakarta Selatan",
	}

	// print object sebelum dijalankan function
	fmt.Println("object sebelum dijalankan function =", reo) // nama 'Mr.' belom ada

	//jalankan function Married pada object reo
	reo.Married()

	// print object setelah dijalankan fucntion
	fmt.Println("object setelah dijalankan function =", reo) // nama 'Mr.' sudah ditambahkan

}

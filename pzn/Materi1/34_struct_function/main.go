package main

import "fmt"

// struct function -> membuat sebuah function seolah-olah milik object struct

// buat struct terlebih dahulu
type Mahasiswa struct {
	ID      int
	Nama    string
	Email   string
	Alamat  string
	Jurusan string
}

// Membuat function seolah-olah milik object Mahasiswa func (object Mahasiswa)
func (inputObject *Mahasiswa) SayHello(inputNama string) {
	fmt.Println("Hello ", inputNama, "My Name is", inputObject.Nama)
}

func main() {
	// buat object dari struct Mahasiswa
	reo := &Mahasiswa{
		ID:      1,
		Nama:    "Reo Sahobby",
		Email:   "reoshby@gmail.com",
		Alamat:  "ragunan",
		Jurusan: "teknik informatika",
	}

	// Panggil function Sayhello dari object
	reo.SayHello("Reo Sahobby")
}

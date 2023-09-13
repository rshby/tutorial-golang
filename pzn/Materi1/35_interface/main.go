package main

import "fmt"

// buat interface
type HasName interface {
	GetName() string
}

// buat function yang menggunakan interface sebagai parameter inputan
func SayHello(inputObject HasName) {
	fmt.Println("Hello", inputObject.GetName())
}

// buat struct Mahasiswa
type Mahasiswa struct {
	ID   int
	Nama string
}

// buat function seolah olah milik object struct Mahasiswa
func (inputObject Mahasiswa) GetName() string {
	return inputObject.Nama
}

// buat struct Animal
type Animal struct {
	Nama  string
	Suara string
}

// buat function seolah olah milih object struct Animal
func (inputObject Animal) GetName() string {
	return inputObject.Nama
}

func main() {
	// buat object dari struct Mahasiswa
	reo := Mahasiswa{
		ID:   1,
		Nama: "Reo Sahobby",
	}

	// panggil function
	SayHello(reo)

	// buat object dari struct Animal
	kucing := Animal{
		Nama:  "Oyen",
		Suara: "Miaw",
	}

	// Panggil function
	SayHello(kucing)
}

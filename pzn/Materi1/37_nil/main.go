package main

import "fmt"

// nil adalah data yang kosong -> biasa digunakan untuk kondisi if

// buat function
func NewMap(inputNama string) map[string]string {
	if inputNama == "" {
		return nil
	} else {
		return map[string]string{
			"Nama": inputNama,
		}
	}
}

func main() {
	// buat variabel
	var person map[string]string
	reo := NewMap("reo")
	budi := NewMap("")

	// buat if pengecekan tanpa kondisi nil
	if person["Nama"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	// buat pengecekan untuk variabel person menggunakan kondisi nil
	if person == nil {
		fmt.Println("Data Person Kosong")
	} else {
		fmt.Println(person)
	}

	// buat pengecekan untuk variabel reo menggunakan kondisi nil
	if reo == nil {
		fmt.Println("Data Reo Kosong")
	} else {
		fmt.Println(reo)
	}

	// buat pengecekan untuk variabel budi menggunakan kondisi nil
	if budi == nil {
		fmt.Println("Data Febi Kosong")
	} else {
		fmt.Println(budi)
	}
}

package main

import (
	"fmt"
	"reflect"
)

// membuat struct
type Karyawan struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// == 1. membuat struct ==
	fmt.Println("- 1. membuat struct -")

	// == 2. membuat object dari struct ==
	fmt.Println("\n- 2. membuat object dari struct -")

	// cara 1 -> membuat struct kosong setelah itu mengisi value ke dalam propertynya
	john := &Karyawan{}
	john.Name = "John Doe"
	john.Age = 22
	john.Address = "jakarta selatan"

	fmt.Println("value object john:", *john)

	// cara 2 -> membuat struct langsung mengisi propertynya (lebih enak)
	michael := &Karyawan{
		Name:    "Michael Kors",
		Age:     25,
		Address: "jakarta barat",
	}
	fmt.Println("value object langsung diisi:", *michael)

	// cara 3 -> mengisi langsung seperti input parameter (harus urut)
	lily := &Karyawan{"Lily Asap", 25, "PIK"}
	fmt.Println("value object langsung di parameter:", *lily)

	// == 3. print object ==
	fmt.Println("\n- 3. print object -")

	// cara 1 -> print berdasarkan property yang dipilih
	fmt.Println("value age object john:", john.Age)
	fmt.Println("value address object michael:", michael.Address)

	// cara 2 -> print semua isi property object
	fmt.Println("-- cara 2 --")
	valueJohn := reflect.ValueOf(*john)
	println("valueJohn numfield:", valueJohn.NumField())
	for i := 0; i < valueJohn.NumField(); i++ {
		fmt.Println("nama property:", valueJohn.Type().Field(i).Name, "value isi:", valueJohn.Field(i))
	}

	// cara 3 -> print menggunakan looping for range jika object dimasukkan ke dalam
	fmt.Println("-- cara 3 --")
	collectionKaryawan := []*Karyawan{john, michael, lily}
	for id, value := range collectionKaryawan {
		fmt.Println("---- id:", id)
		valueObject := reflect.ValueOf(*value)
		for i := 0; i < valueObject.NumField(); i++ {
			fmt.Println("property:", valueObject.Type().Field(i).Name, "valuenya:", valueObject.Field(i))
		}
	}
}

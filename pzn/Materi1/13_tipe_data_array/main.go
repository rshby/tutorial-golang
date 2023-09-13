package main

import (
	"fmt"
	"reflect"
)

// tipe data array pada golang
func main() {
	// == 1. membuat array kosong ==
	fmt.Println("- 1. membuat array kosong")

	var bulan [12]string
	var nilaiUjian [3]int

	fmt.Println("variabel bulan dengan tipe data :", reflect.TypeOf(bulan))
	fmt.Println("variabel nilaiUjian dengan tipe data :", reflect.TypeOf(nilaiUjian))

	// == 2. mengisi array dengan value ==
	fmt.Println()
	fmt.Println("- 2. mengisi array dengan value -")

	// mengisi nama bulan
	bulan[0] = "Januari"
	bulan[1] = "Februari"
	bulan[2] = "Maret"

	// mengisi value nilaiUjian
	nilaiUjian[0] = 80
	nilaiUjian[1] = 90
	nilaiUjian[2] = 100

	// == 3. print array sesuai dengan index yang dipilih ==
	fmt.Println()
	fmt.Println("- 3. print array sesuai dengan yang dipilih -")
	fmt.Println("value array bulan index ke-0 adalah :", bulan[0], "dengan tipe data", reflect.TypeOf(bulan[0]))
	fmt.Println("value array nilaiUjian index ke-1 adalah :", nilaiUjian[1], "dengan tipe data", reflect.TypeOf(nilaiUjian[1]))

	// == 4. print semua array menggunakan for range
	fmt.Println()
	fmt.Println("- 4. print semua array menggunakan for range -")
	for idx, item := range nilaiUjian {
		fmt.Println("index ke", idx, "berisi value :", item, "dengan tipe data", reflect.TypeOf(item))
	}

	// == 5. membuat array langsung mengisi dengan value ==
	myCollection := [2]string{
		"Nike Air Force 1", "New Balance SC Elite V3",
	}

	myNumber := [3]int{1, 2, 3}

	fmt.Println()
	fmt.Println("- 5. membuat array langsung mengisi dengan value -")
	for _, item := range myCollection {
		fmt.Printf("%s | ", item)
	}

	fmt.Println()
	for _, item := range myNumber {
		fmt.Printf("number : %d | ", item)
	}

	// == 6. mengetahui panjang array dengan len ==
	fmt.Println()
	fmt.Println("- 6. mengetahui panjang array dengan len() -")

	panjangArray := len(bulan)
	fmt.Println("panjang array bulan adalah :", panjangArray)
	fmt.Println("capacity array bulan adalah :", cap(bulan))
}

package main

import (
	"fmt"
	"reflect"
)

// operasi matematika di golang
func main() {
	// == 1. penjumlahan ==
	x := 10
	y := 5
	z := x + y
	fmt.Println("- 1. penjumlahan -")
	fmt.Println("value dari variabel z adalah :", z, ". dengan tipe data :", reflect.TypeOf(z))

	// penjumlahan float + int -> hasilnya selalu float
	hasilPenjumalah := 10.5 + 2
	fmt.Println("value dari penjumlahan float + int adalah :", hasilPenjumalah, ". dengan tipe data", reflect.TypeOf(hasilPenjumalah))

	// == 2. pengurangan ==
	x = 80
	y = 20
	z = x - y
	fmt.Println()
	fmt.Println("- 2. pengurangan -")
	fmt.Println("value pengurangan int-int adalah :", z, ". dengan tipe data", reflect.TypeOf(z))

	// pengurangan int - float -> hasilnya selalu float
	hasilPengurangan := 100 - 50.5
	fmt.Println("value pengurangan int-float adalah :", hasilPengurangan, ". dengan tipe data", reflect.TypeOf(hasilPengurangan))

	// pengurangan float - int -> hasilnya selalu float
	hasilPengurangan = 100.5 - 10
	fmt.Println("value pengurangan float-int adalah :", hasilPengurangan, ". dengan tipe data", reflect.TypeOf(hasilPengurangan))

	// == 3. perkalian ==
	// perkalian int * int -> hasilnya int
	hasilPerkalianInt := 10 * 10
	fmt.Println()
	fmt.Println("- 3. perkalian -")
	fmt.Println("value dari perkalian int*int adalah :", hasilPerkalianInt, ". dengan tipe data", reflect.TypeOf(hasilPerkalianInt))

	// perkalian int * float -> hasilnya float
	hasilPerkalian := 100 * 1.5
	fmt.Println("value perkalian (int * float) adalah :", hasilPerkalian, ". dengan tipe data", reflect.TypeOf(hasilPerkalian))

	// perkalian float * int -> hasilnya float
	hasilPerkalian = 12.5 * 2
	fmt.Println("value perkalian (float * int) adalah :", hasilPerkalian, ". dengan tipe data", reflect.TypeOf(hasilPerkalian))

	// perkalian float * float
	hasilPerkalian = 12.4 * 3.2
	fmt.Println("value perkalian (float * float) adalah :", hasilPerkalian, ". dengan tipe data", reflect.TypeOf(hasilPerkalian))

	// == 4. pembagian ==
	// pembagian int/int
	hasilPembagianInt := 100 / 10
	fmt.Println()
	fmt.Println("- 4. pembagian -")
	fmt.Println("hasil pembagian (int / int) adalah :", hasilPembagianInt, ". dengan tipe data", reflect.TypeOf(hasilPembagianInt))

	// pembagian int / float -> hasilnya float
	hasilPembagian := 100 / 10.2
	fmt.Println("hasil pembagian (int / float) adalah :", hasilPembagian, ". dengan tipe data", reflect.TypeOf(hasilPembagian))

	// pembagian (float / int) -> hasilnya float
	hasilPembagian = 100.5 / 3
	fmt.Println("hasil pembagian (float / int) adalah :", hasilPembagian, "dengan tipe data", reflect.TypeOf(hasilPembagian))

	// pembagian (float / float) -> hasilnya float
	hasilPembagian = 120.45 / 3.33
	fmt.Println("hasil pembagian (float / float) adalah :", hasilPembagian, "dengan tipe data", reflect.TypeOf(hasilPembagian))

	// == 5. augmented assignment ==
	nilai := 10
	nilai += 10
	fmt.Println()
	fmt.Println("- 5. augmented assignment -")
	fmt.Println("hasil dari (nilai += 10) adalah :", nilai, "dengan tipe data", reflect.TypeOf(nilai))

	nilai -= 3
	fmt.Println("hasil dari (nilai -= 3) adalah :", nilai, "dengan tipe data", reflect.TypeOf(nilai))

	nilai *= 2
	fmt.Println("hasil dari (nilai *= 2) adalah :", nilai, "dengan tipe data", reflect.TypeOf(nilai))

	nilai /= 2
	fmt.Println("hasil dari (nilai /= 2) adalah :", nilai, "dengan tipe data", reflect.TypeOf(nilai))

	// == 6. unary operator ==
	myValue := 10.5
	myValue++ // ditambah 1.0
	fmt.Println()
	fmt.Println("- 6. unary operator -")
	fmt.Println("hasil dari (myValue++) adalah :", myValue, "dengan tipe data", reflect.TypeOf(myValue))

	myValue-- // dikurangi 1.0
	fmt.Println("hasil dari (myValue--) adalah :", myValue)

	// == 7. angka positif dan negatif ==
	angkaPositif := 200
	angkaNegatif := -50
	fmt.Println()
	fmt.Println("- 7. angka positif dan negatif -")

	operasiMatematika := angkaPositif + angkaNegatif
	fmt.Println("hasil dari (positif + negatif) adalah :", operasiMatematika, "dengan tipe data", reflect.TypeOf(operasiMatematika))

	operasiMatematika = angkaNegatif + angkaPositif
	fmt.Println("hasil dari (negatif + positif) adalah :", operasiMatematika, "dengan tipe data", reflect.TypeOf(operasiMatematika))

	operasiMatematika = angkaPositif - angkaNegatif
	fmt.Println("hasil dari (positif - negatif) adalah :", operasiMatematika, "dengan tipe data", reflect.TypeOf(operasiMatematika))

	operasiMatematika = angkaNegatif - angkaPositif
	fmt.Println("hasil dari (negatif - positif) adalah :", operasiMatematika, "dengan tipe data", reflect.TypeOf(operasiMatematika))
}

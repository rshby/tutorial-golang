package main

import (
	"fmt"
	"reflect"
)

// operasi perbandingan
func main() {
	// == 1. sama dengan
	// perbandingan string dengan string
	nama1 := "reo"
	nama2 := "reo"
	result := nama1 == nama2
	fmt.Println("- 1. sama dengan -")
	fmt.Println("hasil perbandingan string nama1 dan string nama2 adalah :", result, "dengan tipe data", reflect.TypeOf(result))

	nama2 = "budi"
	result = nama1 == nama2
	fmt.Println("hasil perbandingan nama1 dan nama2 adalah :", result, "dengan tipe data", reflect.TypeOf(result))

	// perbandingan int dengan int
	hasilPerbandingan := 10 == 10
	fmt.Println("hasil perbandingan int dengan int adalah :", hasilPerbandingan, "dengan tipe data", reflect.TypeOf(hasilPerbandingan))

	// hasil perbandingan int dengan float
	hasilPerbandingan = 10 == 10.5
	fmt.Println("hasil perbandingan int dengan float adalah :", hasilPerbandingan, "dengan tipe data", reflect.TypeOf(hasilPerbandingan))

	// == 2. tidak sama dengan ==
	// perbandingan string != string
	hasilPerbandingan = "reo" != "andi"
	println()
	fmt.Println("- 2. tidak sama dengan")
	fmt.Println("hasil dari ('reo' != 'andi') adalah :", hasilPerbandingan, "dengan tipe data", reflect.TypeOf(hasilPerbandingan))

	// perbandingan int != int
	hasilPerbandingan = 10 != 10
	fmt.Println("hasil perbandingan (10 != 10) adalah :", hasilPerbandingan, "dengan tipe data", reflect.TypeOf(hasilPerbandingan))

	// perbandingan int != float
	hasilPerbandingan = 10 != 15.5
	fmt.Println("hasil perbandingan (10 != 15.5) adalah :", hasilPerbandingan, "dengan tipe data", reflect.TypeOf(hasilPerbandingan))

	// == 3. lebih besar ==
	// perbandingan string > string
	hasilPerbandingan = "halo" > "reo"
	fmt.Println()
	fmt.Println("- 3. lebih besar -")
	fmt.Println("hasil perbandingan ('halo' > 'reo') adalah :", hasilPerbandingan, "dengan tipe data", reflect.TypeOf(hasilPerbandingan))

	// perbandingan int > int
	hasilPerbandingan = 10 > 5
	fmt.Println(fmt.Sprintf("hasil perbandingan (%v > %v) adalah %v, dengan tipe data %v", 10, 5, hasilPerbandingan, reflect.TypeOf(hasilPerbandingan)))

	// hasil perbandingan int > float
	varInt := 10
	varFloat := 2.5
	hasilPerbandingan = float64(varInt) > varFloat
	fmt.Println(fmt.Sprintf("hasil perbandingan (%v > %v) adalah %v, dengan tipe data %v", float64(varInt), varFloat, hasilPerbandingan, reflect.TypeOf(hasilPerbandingan)))

	// == 4. lebih besar sama dengan ==
	hasilPerbandingan = 10 >= 10
	fmt.Println()
	fmt.Println("- 4. lebih besar sama dengan -")
	fmt.Println("hasil perbandingan (10 >= 10) adalah :", hasilPerbandingan, "dengan tipe data", reflect.TypeOf(hasilPerbandingan))

	// == 5. lebih kecil ==
	hasilPerbandingan = 4 < 4
	fmt.Println()
	fmt.Println("- 5. lebih kecil -")
	fmt.Println("hasil perbandingan (4 < 4) adalah :", hasilPerbandingan, "dengan tipe data", reflect.TypeOf(hasilPerbandingan))

	// == 6. lebih kecil sama dengan ==
	hasilPerbandingan = 4 <= 4
	fmt.Println()
	fmt.Println("- 6. lebih kecil sama dengan -")
	fmt.Println("hasil perbandingan (4 <= 4) adalah :", hasilPerbandingan, "dengan tipe data", reflect.TypeOf(hasilPerbandingan))
}

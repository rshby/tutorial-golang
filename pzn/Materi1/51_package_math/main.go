package main

import (
	"fmt"
	"math"
)

func main() {
	// Round() -> membulatkan float64 ke atas atau bawah sesuai dengan yg paling dekat
	fmt.Println(math.Round(10.7))

	// Floor() -> memaksa membulatkan float64 ke bawah
	fmt.Println(math.Floor(5.7))

	// Ceil() -> memaksa membulatkan float64 ke atas
	fmt.Println(math.Ceil(10.3))

	// Max(..., ...) -> mengembalikan nilai float64 yg paling besar
	fmt.Println(math.Max(10.6, 8.7))

	// Min(..., ...) -> mengembalikan nilai float64 yg paling kecil
	fmt.Println(math.Min(9.7, 2.0))
}

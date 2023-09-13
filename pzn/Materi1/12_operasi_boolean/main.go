package main

import (
	"fmt"
	"reflect"
)

// operasi boolean pada golang
func main() {
	// == 1. lebih besar ==
	nilaiUTS := 90
	nilaiUAS := 85
	batasStandar := 80

	fmt.Println("- 1. lebih besar -")
	lulus := nilaiUAS > batasStandar
	fmt.Println("apakah lulus (nilaiUAS > batasStandar) :", lulus, "dengan tipe data", reflect.TypeOf(lulus))

	// == 2. lebih kecil ==
	tidakLulus := (nilaiUAS+nilaiUTS)/2 <= batasStandar
	fmt.Println()
	fmt.Println("- 2. lebih kecil -")
	fmt.Println("apakah tidak lulus (rata-rata <= batasStandar) :", tidakLulus, "ternyata", func(tidakLulus bool) string {
		if tidakLulus {
			return "tidak lulus"
		} else {
			return "lulus"
		}
	}(tidakLulus))

	// == 3. And && ==
	lapar := true
	capek := true

	butuhMakan := lapar && capek
	fmt.Println()
	fmt.Println("- 3. And && -")
	fmt.Println("apakah saya butuh makan :", butuhMakan)

	// == 4. OR || ==
	fmt.Println()
	fmt.Println("- 4. OR || -")

	tidakSibuk := false
	bolehIstirahat := tidakSibuk || capek
	fmt.Println("apakah saya boleh istirahat :", bolehIstirahat, "karena", func(tidaksibuk, capek bool) string {
		if tidaksibuk {
			return "saya sudah tidak sibuk"
		} else {
			return "saya capek"
		}
	}(tidakSibuk, capek))
}

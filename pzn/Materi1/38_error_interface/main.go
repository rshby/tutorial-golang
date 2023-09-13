package main

import (
	"errors"
	"fmt"
)

// buat function pembagian
func pembagian(inputAngka, inputPembagi int) (int, error) {
	if inputPembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		hasil := inputAngka / inputPembagi
		return hasil, nil
	}
}

func main() {
	// panggil function
	hasil, e := pembagian(10, 0)
	if e != nil {
		fmt.Println("Kesalahan pembagian =", e.Error())
	} else {
		fmt.Println("Hasil pembagian adalah =", hasil)
	}
}

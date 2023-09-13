package main

import (
	"fmt"
	"strconv"
)

/*
package strconv
- untuk melakukan konversi ke tipe data yang berbeda
*/

func main() {
	// ParseBool() -> merubah dari string ke boolean
	if boolean, err := strconv.ParseBool("true"); err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(boolean)
	}

	// ParseFloat() -> merubah dari string ke float
	if number, err := strconv.ParseFloat("5.5", 10); err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(number)
	}

	// ParseInt() -> merubah dari string ke int64
	if number, err := strconv.ParseInt("1000000", 10, 64); err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(number)
	}

	// FormatBool() -> merubah dari boolean ke string
	fmt.Println(strconv.FormatBool(true))

	// FormatFloat() -> merubah dari float ke string
	fmt.Println(strconv.FormatFloat(8.8, 'f', -1, 64))

	// FormatInt() -> merubah dari int ke string
	fmt.Println(strconv.FormatInt(100, 10))

	// Itoa() -> merubah int ke string
	fmt.Print(strconv.Itoa(60))

	// Atoi() -> merubah dari string ke int
	if number, err := strconv.Atoi("350"); err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(number)
	}
}

package main

import (
	"fmt"
	"regexp"
)

/*
berisi utilitas untuk melakukan pencarian regular expression
regex di golang menggunakan library C yg dibuat oleh google bernama RE2
*/

func main() {
	// regexp.MustCompile(string) -> membuat regex depannya e belakang o huruf kecil semua
	var regex = regexp.MustCompile(`e([a-z])o`)

	// regexp.MatchString(string) bool -> mengecek apakah regex match dengan string
	fmt.Println(regex.MatchString("reo"))
	fmt.Println(regex.MatchString("rea"))
	fmt.Println(regex.MatchString("eko"))

	// regexp.FindAllString(string, max) -> mencari string yg match dengan maximum jumlah hasil
	fmt.Println(regex.FindAllString("eko edo eki ero", 10))
}

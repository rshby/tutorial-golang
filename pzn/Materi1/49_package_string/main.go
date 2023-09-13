package main

import (
	"fmt"
	"strings"
)

func main() {
	//Trim() -> menghapus karakter di awal dan akhir dalam string (menghapus * di awal)
	fmt.Println(strings.Trim("*****Reo*sahobby*****", "*"))

	// ToTitle() -> membuat text menjadi besar semua sama kayak ToUpper()
	fmt.Println(strings.ToTitle("budi siswanto"))

	// ToLower() -> membuat menjadi huruf kecil semua
	fmt.Println(strings.ToLower("REO"))

	// ToUpper() -> membuat menjadi huruf besar semua
	fmt.Println(strings.ToUpper("sugeng riyadi"))

	// Split() -> memotong string berdasarkan separator menjadi slice
	textSplit := strings.Split("text1 text2", " ")
	fmt.Println(textSplit)

	// Contains() -> mengecek apakah ada "reo" di dalam sebuah string "reo sahobby"
	fmt.Println(strings.Contains("reo sahobby", "reo"))

	// ReplaceAll() -> mengganti text "reo" menjadi text "ucup"
	fmt.Println(strings.ReplaceAll("reo reo sahobby", "reo", "ucup"))
}

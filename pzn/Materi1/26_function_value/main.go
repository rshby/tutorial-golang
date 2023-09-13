package main

import "fmt"

// create function
func GetAll() int {
	return 10
}

// create function
func GetAllSum(inputNumber int) int {
	return inputNumber * 2
}

func main() {
	// == 1. membuat function ==
	fmt.Println("- 1. membuat function -")

	// == 2. membuat variabel dengan value sebuah function ==
	fmt.Println("\n- 2. membuat variabel dengan value sebuah function -")

	getAllData := GetAll
	getAllKaliDua := GetAllSum

	// == 3. memanggil/menggunakan vairbel function ==
	fmt.Println("\n- 3. call variabel function -")
	result := getAllData()
	fmt.Println("result:", result)

	result = getAllKaliDua(10)
	fmt.Println("result kali dua:", result)
}

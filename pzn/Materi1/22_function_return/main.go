package main

import "fmt"

// create function dengan return value
func SayHello(inputName string) string {
	return fmt.Sprintf("hallos %v", inputName)
}

// create function with return
func GetGrade(inputNilai int) string {
	if inputNilai >= 85 {
		return "A"
	} else if inputNilai >= 75 {
		return "B"
	} else if inputNilai >= 70 {
		return "C"
	} else {
		return "D"
	}
}

func main() {
	// == 1. create function ==

	// == 2. call function with return ==
	fmt.Println("- 2. call function with return value -")
	greeting := SayHello("reo")

	fmt.Println("hasil function:", greeting)

	nilai := 80
	myGrade := GetGrade(nilai)
	fmt.Println("grades:", myGrade)
}

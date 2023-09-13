package main

import (
	"fmt"
)

// recoursive function -> function yang memanggil dirinya sendiri
// contoh kasus = faktorial

// create function faktorial menggunakan for loop
func FaktorialLoop(inputNumber int) int {
	resultSum := inputNumber
	for i := inputNumber; i >= 1; i-- {
		if i == 1 {
			resultSum *= 1
			break
		}
		resultSum *= (i - 1)
	}
	return resultSum
}

// create function faktorial menggunakan recursive
func FaktorialRekursif(inputNumber int) int {
	if inputNumber == 1 {
		return inputNumber
	}
	return inputNumber * FaktorialRekursif(inputNumber-1)
}

func main() {
	// == 1. call function faktorial menggunakan looping ==
	fmt.Println("- 1. call function faktorial menggunakan looping -")
	result := FaktorialLoop(5)
	fmt.Println("hasil faktorial:", result)

	result = FaktorialRekursif(5)
	fmt.Println("result function rekursif:", result)
}

package main

import "fmt"

// create variadic input parameter function
func GreetingAndSumAll(inputNama string, inputNumber ...int) string {
	sum := 0
	for _, value := range inputNumber {
		sum += value
	}

	response := fmt.Sprintf("hello %v, result sum: %v", inputNama, sum)
	return response
}

func main() {
	// == 1. create function with variadic input parameter ==

	// == 2. call variadic function ==
	fmt.Println("- 2. call variadic function -")

	// call menggunakan input slice
	myName := "reo sahobby"
	result := GreetingAndSumAll(myName, 1, 2, 3, 4)

	fmt.Println(result)

}

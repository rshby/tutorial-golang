package main

import "fmt"

// create function
func SayHelloWithFilter(inputNama string, inputFilter func(string) string) string {
	result := fmt.Sprintf("halo %v", inputFilter(inputNama))
	return result
}

// create function filter
func FilterKata(inputKata string) string {
	if inputKata == "anjing" {
		return "*#@#*"
	} else {
		return inputKata
	}
}

func main() {
	// == 1. create function ==
	fmt.Println("- 1. create function -")

	// == 2. call function ==
	fmt.Println("\n- 2. call function -")
	result := SayHelloWithFilter("reo", FilterKata)
	fmt.Println("result:", result)

	result = SayHelloWithFilter("budi", func(s string) string {
		response := fmt.Sprintf("hehe %v", s)
		return response
	})

	fmt.Println("result langsung dibuat function:", result)
}

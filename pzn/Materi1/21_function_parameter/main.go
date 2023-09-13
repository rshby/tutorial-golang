package main

import "fmt"

// func dengan parameter
func main() {
	// == 1. create function with input parameter ==

	// == 2. call function dengan parameter
	fmt.Println("- 2. call function dengan parameter -")
	fmt.Println("   menggunakan parameter variabel")
	firstName := "reo"
	lastName := "sahobby"
	Greeting(firstName, lastName)
	SayHello(firstName, lastName)
}

// create function dengan input parameter
func Greeting(firstName string, lastName string) {
	fmt.Println("good morning", firstName, lastName)
}

// create function dengan parameter tipe data sama
func SayHello(firstName, lastName string) {
	fmt.Println("haloo", firstName, lastName)
}

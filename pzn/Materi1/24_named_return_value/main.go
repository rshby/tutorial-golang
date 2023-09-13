package main

import "fmt"

// create function with named return
func SayHello() (firsName string, lastName string, message string) {
	firsName = "reo"
	lastName = "sahobby"
	message = fmt.Sprintf("hallo %v %v", firsName, lastName)
	return firsName, lastName, message
}

func main() {
	// == 1. create function with named return ==

	// == 2. call function with named return ==
	fmt.Println("- 2. call function with named return -")
	x, y, z := SayHello()
	fmt.Println("first name:", x)
	fmt.Println("last name:", y)
	fmt.Println("message:", z)
}

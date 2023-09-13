package main

import "fmt"

func main() {
	// == 1. membuat constant variabel ==
	// wajib diisi dengan valuenya
	const phi = 3.14

	fmt.Println(fmt.Sprintf("nilai dari phi adalah : %v", phi))

	// == 2. membuat multiple constant variable secara lansung ==
	const (
		// wajib langsung diisi dengan valuenya
		firstName = "Reo"
		lastName  = "Sahobby"
		fullName  = firstName + " " + lastName
	)

	fmt.Println(fmt.Sprintf("nama lengkap saya adalah : %v", fullName))
}

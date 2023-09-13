package helper

import "fmt"

func SayHello(inputNama string) {
	fmt.Println("Hello " + inputNama)
}

// huruf depan kecil -> tidak bisa diakses di luar package
func sayGoodBye(inputNama string) {
	fmt.Println("Good Bye to ", inputNama)
}

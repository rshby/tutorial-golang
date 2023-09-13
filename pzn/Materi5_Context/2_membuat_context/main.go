package main

import (
	"context"
	"fmt"
)

/*
=== Membuat Context ===
karena context adalah sebuah interface, untuk membuat context kita butuh sebuah struct yang sesuai dengan kontrak interface Context
Namun kita tidak perlu membuatnya secara manual
di Golang package Context terdapat function yang bisa kita gunakan untuk membuat Context
*/

/*
Function membuat Context
context.Background() -> membuat context kosong. tidak pernah dibatalkan, tidak memiliki timeout, dan tidak memiliki value apapun. Biasanya digunakan di main function atau dalam test, atau dalam awal proses request terjadi

context.TODO() -> membuat context kosong seperti Background(), namun biasanya menggunakan ini ketika belum jelas context apa yang ingin digunakan
*/

func main() {
	// membuat context background
	background := context.Background()
	fmt.Println(background)
	fmt.Println(background.Done())

	// membuat context TODO
	todo := context.TODO()
	fmt.Println(todo)
}

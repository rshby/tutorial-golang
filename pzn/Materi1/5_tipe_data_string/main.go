// Tipe Data String
package main

import "fmt"

func main() {
	nama := "reo sahobby"

	fmt.Println("Starting")
	fmt.Println("reo")
	fmt.Println("reo sahobby")
	fmt.Println("reo sahobby logitect samsung")

	fmt.Println("panjang dari", nama, "= ", len(nama))
	fmt.Println("reo"[0]) // mengambil huruf index ke 0 -> ditampilkan kode byte hurufnya
	fmt.Println("reo"[1])
}

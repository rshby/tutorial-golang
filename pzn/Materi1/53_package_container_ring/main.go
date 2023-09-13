package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
container ring -> adalah immplementasi struktur data circular list
circular list -> struktur data ring, dimana akhir element akan kembali ke awal element
*/

func main() {
	// membuat variabel circular list
	data := ring.New(5)

	// memasukkan data ke ring menggunakan for
	for i := 0; i < data.Len(); i++ {
		data.Value = "value " + strconv.Itoa(i)
		data = data.Next()
	}

	// menampilkan isi variabel link menggunakan Do(func(){})
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}

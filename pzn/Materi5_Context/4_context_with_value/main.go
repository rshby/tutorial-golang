package main

import (
	"context"
	"fmt"
)

/*
=== Context with Value ===
-> pada saat awal kita membuat context, context tidak memiliki value
-> kita bisa menambahkan sebuah value dengan datanya berupa Pair(key - value) ke dalam context
-> saat kita menambakan value ke dalam context, maka secara otomatis akan tercipta child baru, artinya original context tidak akan berubah sama sekali
-> untuk membuat nemabahkan value ke context kita bisa menggunakan function context.WithValue(parent, key, value)
*/

func main() {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// === Get Value Context
	fmt.Println(contextF.Value("f")) // menampilkan value f
	fmt.Println(contextF.Value("c")) // menampilkan value parrentnya
	fmt.Println(contextF.Value("b")) // tidak bisa, karena beda parrent
	fmt.Println(contextA.Value("b")) // tidak bisa mengambil value child

}

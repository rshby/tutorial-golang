package main

import (
	"container/list"
	"fmt"
)

/*
adalah implementasi struktur data double linked list di golang
*/

func main() {
	data := list.New()
	data.PushBack("reo")
	data.PushBack("sahobby")
	data.PushBack("budi")
	data.PushBack("siswanta")
	data.PushBack("nugraha")

	// menampilkan data -> Next() untuk data selanjutnya
	fmt.Println(data.Front().Value)
	fmt.Println(data.Front().Next().Value)

	// menampilkan data -> Prev() untuk data sebelumnya
	fmt.Println(data.Back().Value)
	fmt.Println(data.Back().Prev().Value)

	// menampilkan semua data menggunakan for -> iterasi maju
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// menampilkan semua data menggunakan for -> iterasi mundur
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}

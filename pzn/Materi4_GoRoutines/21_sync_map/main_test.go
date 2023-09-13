package main

import (
	"fmt"
	"sync"
	"testing"
)

/*
Golang memiliki sebuah struct bernama sync.Map
- Map mirip dengan dengan map pada golang, namun yang membedakan Map ini aman untuk menggunakan concurrent menggunakan goroutine
Ada beberapa function yang bisa kita gunakan di Map:
-- Store(key, value) -> menyimpan data ke Map
-- Load(key) -> mengambil data dari Map menggunakan Key
-- Delete(key) -> menghapus data di Map menggunakan Key
-- Range(function(key, value)) -> melakukan iterasi seluruh data di Map
*/

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMain(m *testing.M) {
	group := &sync.WaitGroup{}
	data := &sync.Map{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

	fmt.Println("Test Complete!!")
}

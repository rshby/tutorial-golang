package main

import (
	"fmt"
	"sync"
	"testing"
)

/*
waitgroup -> fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutines tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
kasus seperti ini bisa menggunakan waitgroup
untuk menandai bahwa ada proses goroutine, kita dapat menggunakan method Add(int), setelah proses goroutine selesai, kita bisa menggunakan method Done()
untuk menunggu semua proses selesai, kita bisa menggunakan method Wait()
*/

// menbuat function run async
func RunAsync(group *sync.WaitGroup, inputNumber int) {
	defer group.Done()

	group.Add(1)

	// tulis proses codingan di bawah ini
	fmt.Println(inputNumber)
}

// buat function test wait
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go RunAsync(group, i)
	}

	group.Wait()
	fmt.Println("Test Complete!")
}

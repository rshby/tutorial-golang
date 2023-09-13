package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*
=== Context with Deadline ===
-> selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa menggunakan deadline
-> pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu dari sekarang, kalo deadline ditentukan kapan waktu timeoutnya, misal jam 12 siang hari ini
-> untuk membuat context dengan deadline kita bisa menggunakan function context.WithDeadline(parent, time)
*/

// buat function yang mengimplementasikan context
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	// buat goroutine
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	return destination
}

// function run program
func main() {
	fmt.Println("Jumlah Goroutine Awal =", runtime.NumGoroutine())

	// buat variabel context
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	// consume function CreateCounter
	destination := CreateCounter(ctx)
	fmt.Println("Jumlah goroutine berjalan =", runtime.NumGoroutine())

	for i := range destination {
		fmt.Println("Counter", i)
	}

	fmt.Println("Jumlah Goroutine Akhir =", runtime.NumGoroutine())
}

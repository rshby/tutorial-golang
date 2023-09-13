package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*
=== Context With Timeout ===
-> selain menambahkan value ke context, dan juga sinyal cancel, kita juga bisa menambahkan sinyak timeout secara otomatis dengan menggunakan pengaturan timeout
-> dengan menggunakan penagturan timeout, kita tidak perlu melakukan eksekusi cancel secara manual, cancel akan otomatis dieksekusi jika timeout sidah terlewati
-> penggunaan context dengan timeout sangat cocok misal kita melakukan query ke database atau API, namun ingin menentukan batas maksimal timeoutnya
-> untuk membuat context dengan timeout, kita dapat menggunakan function context.WithTimeout(parent, duration)
*/

// buat function yang mengimplementasikan context timeout
func CreateCounter(ctx context.Context) chan int {
	result := make(chan int)

	// buat goroutine
	go func() {
		defer close(result)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				result <- counter
				counter++
				time.Sleep(500 * time.Millisecond) // simulasi slow process 0.5 second
			}
		}
	}()
	return result
}

// func run program
func main() {
	fmt.Println("Jumlah goroutine awal =", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	// consume function CreateCounter
	destination := CreateCounter(ctx)

	fmt.Println("Jumlah goroutine berjalan =", runtime.NumGoroutine())
	for i := range destination {
		fmt.Println("Counter ke", i)
	}

	fmt.Println("Jumlah goroutine akhir =", runtime.NumGoroutine())
}

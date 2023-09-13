package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*
=== Context With Cancel ===
-> selain menambahkan value ke context, kita juga bisa menambahkan sinyal cancel ke context
-> kapan sinyal cancel diperlukan dalam context?
-> biasanya ketika kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke proses tersebut
-> biasanya proses ini berupa goroutine yang berbeda, sehingga dengan mudah jika kita ingin mmbatalkan eksekusi goroutine, kita bisa mengirim sinyal cancel ke contextnya
-> namun ingat goroutine yang menggunakan context tetap harus melakukan pengecekan terhadap contexnya, jika tidak maka tidak ada gunanya
-> untuk membuat context dengan cancel signal, kita bisa menggunakan function context.WithCancel(parent)
*/

// buat function -> contoh context leak
func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		// perulangan tidak berhenti dan selalu mengirim ke channel
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

// buat function yang benar
func CreateCounterBenar(ctx context.Context) chan int {
	destination := make(chan int)

	// buat goroutine
	go func() {
		defer close(destination)
		counter := 1

		// buat perulangan tanpa henti
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func main() {
	fmt.Println(runtime.NumGoroutine())

	// consume function CreateCounter
	destination := CreateCounter()

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine =", runtime.NumGoroutine())
	fmt.Println()

	fmt.Println("Jumlah Goroutine Awal =", runtime.NumGoroutine())

	// buat parent channelnya terlebih dahulu
	parent := context.Background()

	// buat context yang berisi sinyal cancel
	ctx, cancel := context.WithCancel(parent)

	// consume function CreateCounterBenar(ctx)
	destinationBenar := CreateCounterBenar(ctx)

	for n := range destinationBenar {
		fmt.Println("Counter Benar", n)
		if n == 10 {
			break
		}
	}
	cancel() // mengirim sinyal cancel ke context
	time.Sleep(2 * time.Second)
	fmt.Println("Jumlah Goroutine Akhir =", runtime.NumGoroutine())
}

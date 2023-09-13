package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
Golang memiliki package yang bernama sync/atomic
Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent
contohnya sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikkan angkan di counter. Hal ini sebenarnya bisa digunakan menggunakan atomic package
*/

func TestAtomic(t *testing.T) {
	group := &sync.WaitGroup{}
	var x int64 = 0
	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				group.Add(1)
				atomic.AddInt64(&x, 1)
				group.Done()
			}
		}()
		group.Wait()
		fmt.Println("counter =", x)
	}

}

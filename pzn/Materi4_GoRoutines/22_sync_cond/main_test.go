package main

import (
	"fmt"
	"sync"
	"testing"
)

/*
COND
Cond -> implementasi locking berbasis kondisi
- Cond membutuhkan Locker(bisa menggunakan Mutex atau RWMutex) untuk implementasi lockingnya, namun berbeda dengan Locker biasanya, di Cond terdapat function Wait() untuk menunggu apakah perlu menunggu atau tidak
- function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan function Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi
- Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)
*/

var cond = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Millisecond)
	// 		cond.Signal()
	// 	}
	// }()

	go func() {
		cond.Broadcast()
	}()

	group.Wait()
	fmt.Println("Test Complete!!!")
}

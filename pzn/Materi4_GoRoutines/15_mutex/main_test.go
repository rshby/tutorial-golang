package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	x := 10

	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Counter x =", x)
}

package main

import (
	"fmt"
	"testing"
	"time"
)

// buat function
func DisplayNumber(number int) {
	fmt.Println("Display -", number)
}

// test function
func TestDisplay(t *testing.T) {
	for i := 0; i <= 1048575; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(30 * time.Second)
}

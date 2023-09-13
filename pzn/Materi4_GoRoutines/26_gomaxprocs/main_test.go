package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU =", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread =", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine =", totalGoRoutine)
}

package main

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func TestRunHelloWorld(t *testing.T) {
	go HelloWorld()
	fmt.Println("Upss!")

	time.Sleep(1 * time.Millisecond)
}

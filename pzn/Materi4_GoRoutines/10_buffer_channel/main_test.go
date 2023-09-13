package main

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 5)
	defer close(channel)

	channel <- "Reo"
	channel <- "Sahobby"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
}

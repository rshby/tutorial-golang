package main

import (
	"fmt"
	"testing"
	"time"
)

// membuat function untuk In -> memasukkan data ke channel
func OnlyIn(channel chan<- string) {
	channel <- "Reo Sahobby"
}

// membuat function untuk Out channel -> mengambil data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Millisecond)
	close(channel)
}

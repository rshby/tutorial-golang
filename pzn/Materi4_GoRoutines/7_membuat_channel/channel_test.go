package main

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(inputChannel chan string) {
	time.Sleep(1 * time.Second)
	inputChannel <- "Reo Sahobby"
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Reo Sahobby"
		fmt.Println("Mengirim data ke channel -> selesai")
	}()

	data := <-channel
	fmt.Println(data)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(1 * time.Second)
}

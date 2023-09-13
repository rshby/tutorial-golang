package main

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(inputChannel chan string) {
	time.Sleep(1 * time.Microsecond)
	inputChannel <- "Reo Sahobby"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data!")
		}
		if counter == 2 {
			break
		}
	}
}

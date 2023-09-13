package main

import (
	"fmt"
	"testing"
	"time"
)

/*
Ticker -> representasi kejadian yang berulang
- ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
- untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
- untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()
*/

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

/*
time.Tick()
- kadang kita tidak butuh data Tickernya, hanya butuh channelnya saja
- jika demikian, kita bisa menggunakan function time.Tick(duration), function ini tidak akan mengembalikan ticker hanya mengembalikan channel timernya saja
*/

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	for tick := range channel {
		fmt.Println(tick)
	}
}

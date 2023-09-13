package main

import (
	"fmt"
	"time"
)

// func yang memakai channel
func MenerimaGoroutine(ch <-chan string) {
	fmt.Println("isi channel adalah : ", <-ch)
	fmt.Println("ini dari goroutine loh!")
	fmt.Println("ini dari goroutine loh!!")
	fmt.Println("ini dari goroutine loh!!!")
	fmt.Println("ini dari goroutine loh!!!!")
}

func main() {
	fmt.Println("== program dimulai ==")

	// make channel
	kata := make(chan string)

	// jalankan goroutine
	go MenerimaGoroutine(kata)

	fmt.Println("tidur dulu 4 detik")
	fmt.Println("....")
	fmt.Println("....")
	time.Sleep(4 * time.Second)

	fmt.Println("ok kirim data ke channel")
	kata <- "haloha"

	fmt.Println("....")
	time.Sleep(5 * time.Second)
	fmt.Println("== program selesai ==")
}

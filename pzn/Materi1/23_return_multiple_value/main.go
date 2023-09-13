package main

import (
	"fmt"
	"time"
)

// create function with multiple return
func SayHello(inputName string) (string, time.Time) {
	greeting := fmt.Sprintf("hello %v", inputName)

	return greeting, time.Now()
}

func main() {
	// == 1. create function with multiple return

	// == 2. call function with multiple return ==
	fmt.Println("- 2. call return with multiple return -")
	greeting, time := SayHello("reo")
	fmt.Println("greeting:", greeting)
	fmt.Println("time at:", time.Format("2006-01-02 15:04:05"))

	// == 3. call function hanya memerlukan salah satu return ==
	fmt.Println("\n- 3. call function hanya mengambil salah satu return -")

	// jika ada return yang tidak dipakai bisa dinamai _ (underscore)
	result, _ := SayHello("budi")
	fmt.Println("value result greeting:", result)
}

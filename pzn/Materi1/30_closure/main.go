package main

import "fmt"

func main() {
	counter := 0
	name := "reo"

	tambahClosure := func() {
		name := "azam"
		fmt.Println("name dalam function =", name)
		fmt.Println("nilai counter =", counter)
		counter++
	}

	tambahClosure()
	tambahClosure()

	fmt.Println("nilai akhir counter =", counter)
	fmt.Println("name akhir adalah =", name)
}

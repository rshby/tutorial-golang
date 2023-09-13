package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// == 1. break ==
	fmt.Println("- 1. Break -")

	// break untuk looping slice (langsung stop semua looping ketika ketemu syntax break)
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println("value mySlice:", mySlice, "len:", len(mySlice), "capacity:", cap(mySlice))
	for _, item := range mySlice {
		if item == 4 {
			break
		}

		fmt.Printf("%d, ", item)
	}

	// break untuk looping map
	fmt.Println()
	myMap := map[string]string{
		"nama":   "reo",
		"umur":   "24",
		"alamat": "ragunan",
	}
	fmt.Println("value myMap:", myMap, "len:", len(myMap))
	for key, value := range myMap {
		if strings.ToLower(key) == "umur" {
			fmt.Println("key:", key, "value:", value)
			break
		}

		fmt.Println("key:", key, "value:", value)
	}

	// break nedted for looop
	for outer := 0; outer < 10; outer++ {
		if outer == 2 {
			fmt.Println("break outer loop")
			break
		}

		fmt.Println("value outer: ", outer)
		for inner := 0; inner < 5; inner++ {
			if inner == 3 {
				fmt.Println("  break inner loop")
				break
			}
			fmt.Println("  value inner:", inner)
		}
	}

	// == 2. Continue ==
	fmt.Println("\n- 2. continue -")

	// continue pada array/slice
	mySlice = []int{5, 2, 3, 8, 4, 1, 10, 13}
	sort.Ints(mySlice)
	fmt.Println("value slice:", mySlice, "len:", len(mySlice), "capacity:", cap(mySlice))
	for i, item := range mySlice {
		if i == len(mySlice)-1 {
			fmt.Printf("%d", item)
			break
		}

		if item == 5 {
			continue
		}

		fmt.Printf("%d, ", item)
	}

	// continue pada map
	myMap = map[string]string{
		"nama":   "reo",
		"umur":   "24",
		"agama":  "islam",
		"alamat": "ragunan",
	}
	fmt.Println("\nvalue myMap:", myMap, "len:", len(myMap))

	for key, value := range myMap {
		if strings.ToLower(key) == strings.ToLower("agama") {
			fmt.Println("agama diskip")
			continue
		}
		fmt.Println("key:", key, ". value:", value)
	}
}

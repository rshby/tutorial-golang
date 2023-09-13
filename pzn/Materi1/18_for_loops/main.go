package main

import "fmt"

// for loop pada golang
func main() {
	// == 1. for biasa ==
	fmt.Println("- 1. for biasa -")
	i := 0
	for i < 5 {
		fmt.Println("for biasa ke-", i)
		i++
	}

	// == 2. for loop menggunakan statement ==
	fmt.Println("\n- 2. for loop menggunakan statement -")

	// looping bisaa
	for i := 0; i < 3; i++ {
		fmt.Println("for statement iterasi ke-", i+1)
	}

	// contoh looping slice
	mySlice := []string{"senin", "selasa", "rabu", "kamis", "jumat"}
	fmt.Printf("nama hari: ")
	for i := 0; i < len(mySlice); i++ {
		if i == len(mySlice)-1 {
			fmt.Printf(mySlice[i])
			break
		}

		fmt.Printf("%s, ", mySlice[i])
	}

	// == 3. looping menggunakan for range ==
	fmt.Println("\n\n- 3. looping menggunakan for range -")

	mySlice = []string{"satu", "dua", "tiga"}
	myMap := map[string]string{
		"nama":   "reo sahobby",
		"umur":   "21",
		"alamat": "ragunan",
	}

	// looping slice dengan index
	for i, item := range mySlice {
		fmt.Println(fmt.Sprintf("index ke-%v value: %v", i, item))
	}

	// looping slice tanpa index
	fmt.Printf("urutan angka: ")
	for _, item := range mySlice {
		fmt.Printf("%s, ", item)
	}

	// looping map dengan key-value
	fmt.Println()
	for key, value := range myMap {
		fmt.Printf("key-%s, value: %s\n", key, value)
	}

	// looping map menampilkan hanya value saja
	fmt.Printf("looping map hanya value saja: ")
	for _, value := range myMap {
		fmt.Printf("%s, ", value)
	}
}

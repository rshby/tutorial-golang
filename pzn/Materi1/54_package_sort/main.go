package main

import (
	"fmt"
	"sort"
)

/*
package sort
- berisi utilitas untuk proses pengurutan
- agar data bisa diurutkan, kita harus mengimplementasikan kontrak di interface Sort.Interface
*/

// membut struct
type User struct {
	Name string
	Age  int
}

// membuat alias
type UserSlice []User

// membuat function Len -> menghitung panjang
func (value UserSlice) Len() int {
	return len(value)
}

// membuat function Less -> mengetahui apakah value index i < index j
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

// membuat function Swap -> mengganti value index i dengan value index j
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	// buat object dari struct Users
	users := []User{
		{"reo", 23},
		{"andi", 22},
		{"joko", 35},
		{"rudi", 31},
	}

	// menampilkan users sebelum diurutkan
	fmt.Println(users)

	// mengurutkan users berdasarkan Age
	sort.Sort(UserSlice(users))

	// menampilkan users setelah diurutkan
	fmt.Println(users)
}

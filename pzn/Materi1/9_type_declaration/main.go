package main

import (
	"fmt"
	"reflect"
)

func main() {
	// type declaration -> membuat variabel menjadi tipe data (merubah nama tipe data sesuai dengan nama yang kita tentukan)

	// membuat noKTP menjadi alias untuk string
	type noKTP string

	// membuat variabel dengan tipe data noKTP
	var noId noKTP = "3310250502990002"

	fmt.Println("- 1. type declaration -")
	fmt.Println("value dari variabel adalah :", noId, ". dengan tipe data :", reflect.TypeOf(noId))

	// == 2. membuat alias untuk tipe data int ==
	type nilaiUAS int

	var nilaiFisika nilaiUAS = 96 // membuat variabel dengan tipe data nilaiUAS
	fmt.Println()
	fmt.Println("- 2. membuat alias untuk tipe data int -")
	fmt.Println("value dari variabel nilaiFisika adalah :", nilaiFisika, ". dengan tipe data :", reflect.TypeOf(nilaiFisika))

	// == 3. membuat alias untuk tipe data boolean
	type isLogin bool

	var cekLogin isLogin = true // membuat variabel dengan tipe data isLogin
	fmt.Println()
	fmt.Println("- 3. membuat alias untuk tipe data boolean -")
	if cekLogin {
		fmt.Println("suksesk login, value :", cekLogin, ". dengan tipe data :", reflect.TypeOf(cekLogin))
	} else {
		fmt.Println("gagal login karena cekLogin bernilai :", cekLogin, ". dengan tipe data :", reflect.TypeOf(cekLogin))
	}

	// == 4. membuat alias untuk tipe data slice
	type shoesCollection []string

	// membuat variabel dengan tipe data shoesCollection
	var myCollection shoesCollection = shoesCollection{"Nike Vaporfly 3", "Adidas Adios Pro 3", "Sauconi Endorphine Elite"}
	fmt.Println()
	fmt.Println("- 4. membuat alias untuk tipe data slice -")
	for _, item := range myCollection {
		fmt.Print(fmt.Sprintf("%v - ", item))
	}
}

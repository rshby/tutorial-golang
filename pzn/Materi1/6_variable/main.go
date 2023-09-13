package main

import "fmt"

func main() {
	// == 1. membuat variabel kosong ==
	var name string

	// mengisi variabel kosong dengan sebuah value
	name = "reo sahobby"
	fmt.Printf("nama saya adalah : %s\n", name)

	// mengganti isi variabel dengan sebuah value yang berbeda
	name = "muhammad reo sahobby"
	fmt.Println(fmt.Sprintf("nama saya yang baru adalah %v", name))

	// == 2. membuat variabel dan langsung mengisi dengan sebuah value ==
	// menggunakan var [nama] = [value]
	var myName = "reo sahobby"
	fmt.Println(fmt.Sprintf("isi dari varibel myName adalah : %v", myName))

	// mengganti value dari sebuah variabel
	myName = "ikhsan nur syahbanu"
	fmt.Println(fmt.Sprintf("nama baru saya adalah : %v", myName))

	// == 3. membuat variabel langsung mengisi dengan sebuah value tanpa menggunakan var ==
	fullName := "Reo Sahobby S.Kom"
	fmt.Println(fmt.Sprintf("nama dan gelar saya adalah : %v", fullName))

	// mengganti value dari variabel yang sudah ada
	fullName = "muhammad achmad affandi setiawan purnomo"
	fmt.Println(fmt.Sprintf("nama lengkap saya berubah menjadi : %v", fullName))

	// == 4. membuat multiple varibel ==
	// membuat multiple variable langsung dengan valuenya
	var (
		firstName = "reo"
		lastName  = "sahobby"
	)

	// membuat multiple variable kosong
	var (
		midddleName string
		grandName   string
	)

	// mengisi variable dengan value
	midddleName = "halo"
	grandName = "hihi"

	fmt.Println(fmt.Sprintf("kemudian nama lengkap saya menjadi : %v %v", firstName, lastName))
	fmt.Println(fmt.Sprintf("ini menjadi : %v %v", midddleName, grandName))
}

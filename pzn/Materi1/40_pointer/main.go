package main

import "fmt"

// buat struct Address
type Address struct {
	City    string
	Provice string
}

func main() {
	// buat object dari struct Address
	address1 := Address{
		City:    "Jakarta Selatan",
		Provice: "DKI Jakarta",
	}

	// buat variabel address2 dari address1
	address2 := &address1

	// Kita Ubah adress2
	address2.City = "Jakarta Utara"

	// buat variiabel address3 pointer mengarah ke address1
	address3 := &address1

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	fmt.Println("")

	// kita ubah adress3 (dimana dia pointer yg mengarah ke address1)
	address3.City = "Klaten"
	fmt.Println("address1 setelah diubah address3 =", address1)
	fmt.Println("address2 setelah diubah address3 =", address2)
	fmt.Println("address3 setelah diubah", address3)

	// buat variabel adress3 yang sekaligus mereplace semua object struct Address (menggunakan *)
	*address3 = Address{
		City:    "Sleman",
		Provice: "DIY",
	}

	fmt.Println("")

	// Ketika diprint semua object dari struct Address akan sama dengan isi address3
	fmt.Println("address1 setelah adress3 dibuat ulang", address1)
	fmt.Println("address2 setelah adress3 dibuat ulang", address2)
	fmt.Println("address2 setelah dibuat ulang", address3)

	// membuat variabel pointer baru menggunakan new (isi data kosong)
	address4 := new(Address)
	address4.City = "Jakarta Selatan"

	fmt.Println("")
	fmt.Println("address4 =", address4)
}

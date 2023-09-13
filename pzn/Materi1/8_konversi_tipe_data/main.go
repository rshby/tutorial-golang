package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// == 1. konversi antar number ==
	fmt.Println("- 1. konversi antar number -")
	var number int = 15
	number8 := int8(number)
	number16 := int16(number8)
	number32 := int32(number16)
	number64 := int64(number32)

	// == 2. cek tipe data variabel menggunakan reflect.TypeOf() ==
	fmt.Println("value variabel number :", number, ". dengan tipe data :", reflect.TypeOf(number))
	fmt.Println("value variabel number8 :", number8, ". dengan tipa data :", reflect.TypeOf(number8))
	fmt.Println("value variabel number16 :", number16, ". dengan tipe data :", reflect.TypeOf(number16))
	fmt.Println("value variabel number32 :", number32, ". dengan tipe data :", reflect.TypeOf(number32))
	fmt.Println("value variabel number64 :", number64, ". dengan tipe data :", reflect.TypeOf(number64))

	// == 3.konversi dari string ke number
	// menggunakan int() -> hasil byte
	myString := "hello"
	myNumber := int(myString[0])
	fmt.Println()
	fmt.Println("- 3. konversi dari string ke number -")
	fmt.Println("value dari variabel myNumber :", myNumber, "dengan tipe data", reflect.TypeOf(myNumber))

	// menggunakan strconv.Atoi(string)
	var err error
	myNumber, err = strconv.Atoi("12")
	if err != nil {
		println("gagal konversi karena ada error :", err.Error())
	} else {
		fmt.Println("success konversi. value variabel myNumber :", myNumber, ". dengan tipe data :", reflect.TypeOf(myNumber))
	}

	// == 4. konversi number ke string ==
	// int ke string
	myString = strconv.Itoa(120)
	fmt.Println()
	fmt.Println("- 4. konversi number ke string -")
	fmt.Println("value dari variabel myString :", myString, ". dengan tipe data :", reflect.TypeOf(myNumber))

	// float ke string -> menggunakan fmt.Sprintf()
	myFloat := 11.6
	myString = fmt.Sprintf("%v", myFloat)
	fmt.Println("value dari variabel mystring :", myString, ". dengan tipe data :", reflect.TypeOf(myString))

	// == 5.konversi dari string ke boolean
	// konversi string "true/false" ke boolean
	var myBool bool
	myBool, err = strconv.ParseBool("false")
	if err != nil {
		fmt.Println("gagal konversi karena ada error :", err.Error())
	} else {
		println("success konversi. value variabel myBool :", myBool, ". dengan tipe data :", reflect.TypeOf(myBool))
	}

	// konversi string angka "1" ke boolean -> hanya bisa angka 0 dan 1
	myBool, err = strconv.ParseBool("1")
	if err != nil {
		fmt.Println("gagal konversi, karena ada error :", err.Error())
	} else {
		fmt.Println("success konversi angka ke boolean :", myBool, ". dengan tipe data :", reflect.TypeOf(myBool))
	}

	// == 6. konversi dari boolean ke string ==
	myString = strconv.FormatBool(false)
	println()
	fmt.Println("- 6. konversi dari boolean ke string -")
	fmt.Println("value dari variabel myString adalah :", myString, ". dengan tipe data :", reflect.TypeOf(myString))

	// == 7. konversi dari number ke bolean
	myBool = 1 != 0
	fmt.Println()
	fmt.Println("- 7. konversi dari number ke boolean -")
	fmt.Println("value dari variable myBool :", myBool, ". dengan tipe data :", reflect.TypeOf(myBool))

	// == 8. konversi dari boolean ke number ==
	// manual menggunakan function
	ConvertToInt := func(inputBool bool) int {
		if inputBool {
			return 1
		} else {
			return 0
		}
	}
	myNumber = ConvertToInt(true)
	fmt.Println()
	fmt.Println("- 8. konversi dari boolean ke number -")
	fmt.Println("hasil value variabel myNumber adalah :", myNumber, ". dengan tipe data :", reflect.TypeOf(myNumber))
}

package helper

/*
t.Fail() -> akan menggagalkan unit test namun tetap melanjutkan eksekusi unit test, namun di akhir ketika selesai maka unit test tersebut dianggap gagal

t.FailNow() -> akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

t.Error(args...) -> seperti melakukan log(print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil funcion t.Fail(), sehingga mengakibatkan unit test dianggap gagal
namun karena hanya memanggil function t.Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai

t.Fatal(args...) -> mirip dengan t.Error() hanya saja setelah melakukan log error, dia akan memanggil t.FailNow(), sehingga mengakibatkan eksekusi unit test berhenti
*/

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// error
		t.Fail()

		fmt.Println("TestHelloWorld Eko Done")
	}
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Khannedy" {
		// error
		t.FailNow()

		fmt.Println("TestHelloWorldKhannedy Done")
	}
}

func TestError(t *testing.T) {
	result := HelloWorld("Reo")
	if result != "Hello Reo" {
		//error
		t.Error("Harusnya Hello Reo")
	}

	fmt.Println("TestError Dieksekusi")
}

func TestFatal(t *testing.T) {
	result := HelloWorld("Reo")
	if result != "Hello Reo" {
		//error
		t.Fatal("Harusnya Hello Reo")
	}

	fmt.Println("TestFatal Dieksekusi")
}

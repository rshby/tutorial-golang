package helper

/*
assert vs require

testyfy menyediakan dua package untuk assertion, yaitu assert dan require
saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil t.Fail(), artinya eksekusi unit test akan tetap dijalankan
sedangkan jika kita menggunakan require, jika pengecekan gagal maka require akan memanggil t.FailNow(), artinya eksekusi unit test tidak akan dilanjutkan
*/

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("Dieksekusi!")
}

func TestRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("Dieksekusi!")
}

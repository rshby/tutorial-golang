package helper

/*
biasanya dalam unit test kadang kita ingin melaukan sesuatu sebelum dan sesudah test dieksekusi

jikalau kode yang kita lakukan sebelum dan sesudah selalu sama antar testnya, maka membuat manual di unit test function adalah hal yang membosankan dan terlalu banyak kode duplikat

untungnya di golang terdapat fitur yang bernama testing.M

fitur ini bernama Main, dimana digunakan untuk mengeksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan before and after di test unit

=== testing.M ===
Untuk  mengatur eksekusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter testing.M

jika terdapat function TestMain tersebut, maka secara otomatis golang akan mengeksekusi function ini tiap kali akan menjalankan unit test di sebuah package

dengan ini kita bisa mengatur before and after test sesuai yang kita mau
*/

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum unit test")

	m.Run() // jalankan semua unit test

	fmt.Println("Setelah dieksekusi!")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tidak bisa dijalankan di windows!")
	}

	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result)
}

func TestAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

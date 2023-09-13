package helper

/*
Sub Test
golang mendukung fitur pembuatan function unit test di dalam function unit test

fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test bahasa pemrograman lain

untuk membuat sub test, kita bisa menggunakan function Run()
*/

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum test")

	m.Run()

	fmt.Println("Setelah Testing")
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result)
	})

	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result)
	})
}

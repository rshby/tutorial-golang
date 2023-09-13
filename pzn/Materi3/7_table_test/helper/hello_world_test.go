package helper

/*
sebelumnya kita telah belajar tentang sub test
jika diperhatikan, sebenarnya dengan subtest kita bisa membuat test secara dinamis
dan fitur sub test ini, digunakan programmer golang untuk membuat test dengan konsep table test
table test -> dimana kita menyediakan data berupa slice yang berisi parameter dan ekspektasi hasil dari unit test
lalu slice tersebut kita iterasi menggunakan sub test
*/

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum test")

	m.Run()

	fmt.Println("Setelah Testing")
}

func TestHelloWorldTableTest(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Helloworld(EKo)",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "HelloWorld(Reo)",
			request:  "Reo",
			expected: "Hello Reo",
		},
		{
			name:     "HelloWorld(Budi)",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

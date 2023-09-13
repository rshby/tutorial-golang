package helper

import "testing"

func TestHelloWord(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// unit test failed
		panic("Rsult is not Eko")
	}
}

package testing

import (
	"Materi3/10_membuat_benchmark/helper"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.HelloWorld("Reo")
	}
}

func BenchmarkHelloWorldLia(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.HelloWorld("Eko")
	}
}

package helper

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unit Test Tidak Bisa Jalan di Windows!")
	}

	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result)
}

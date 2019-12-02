package password

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	password := New("123456")
	fmt.Println("len:", len(password), "\n", "password:", password)
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("123456")
	}
}

package math

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4, 6))
	// Output: 10
}

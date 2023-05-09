package ch11

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2}, {1, 2, 4}, {2, 0, 4},
	}

	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("expect: %d, actual: %d", value.out, re)
		}
	}
}

func BenchmarkAdd(bb *testing.B) {
	a := 123
	b := 456
	c := 579
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Printf("expect %d, actual %d", c, actual)
		}
	}
}

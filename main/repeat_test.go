package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("print a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("got %q want %q", repeated, expected)

		}
	})

	t.Run("print a 5 times", func(t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"

		if repeated != expected {
			t.Errorf("got %q want %q", repeated, expected)
		}
	})
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 3))
	//Output: aaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

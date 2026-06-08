package interation

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat test 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("Got %s, expected %s", repeated, expected)
		}
	})
	t.Run("repeat test 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("Got %s, expected %s", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	output := Repeat("a", 5)
	fmt.Println(output)
	// Output: aaaaa
}

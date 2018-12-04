package tmp

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var fibTests = []struct {
		in       string // input
		in2      int    // input
		expected int    // expected result
	}{
		{"((1)23(45))(aB)", 0, 10},
		{"((1)23(45))(aB)", 1, 3},
		{"((1)23(45))(aB)", 2, -1},
		{"((1)23(45))(aB)", 6, 9},
		{"((1)23(45))(aB)", 11, 14},
		{"((>)|?(*'))(yZ)", 11, 14},
	}
	for _, tt := range fibTests {
		actual, _ := Solution(tt.in, tt.in2)
		if actual != tt.expected {
			t.Errorf("Solution(arr []int) = %d; expected %d", actual, tt.expected)
		}
	}
}

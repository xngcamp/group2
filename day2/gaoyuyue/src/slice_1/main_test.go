package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		str string
		idx int
		end int
	}{
		{"((1)23(45))(aB)", 0, 10},
		{"((1)23(45))(aB)", 1, 3},
		{"((1)23(45))(aB)", 2, -1},
		{"((1)23(45))(aB)", 6, 9},
		{"((1)23(45))(aB)", 11, 14},
		{"((>)|?(*'))(yZ)", 11, 14},
	}

	for _, test := range tests {
		if end, _ := Solution(test.str, test.idx); end != test.end {
			t.Errorf("Error: Solution input str: %v idx: %v expect: %v actual: %v", test.str, test.idx, test.end, end)
		}
	}
}

package main

import "testing"

func TestCalMaxUniqueStr(t *testing.T) {
	tests := []struct {
		str    string
		result string
	}{
		{"ababab", "ab"},
		{"abcdefbgac", "cdefbga"},
	}

	for _, test := range tests {
		if result := CalMaxUniqueStr(test.str); result != test.result {
			t.Errorf("Error: CalMaxUniqueStr input str: %v expect: %v actual: %v", test.str, test.result, result)
		}
	}
}

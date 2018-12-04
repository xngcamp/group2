package isbn

import (
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	for _, test := range testCases {
		observed := IsValidISBN(test.isbn)
		if observed != test.expected {
			t.Errorf("FAIL: %s\nIsValidISBN(%q)\nExpected: %t, Actual: %t",
				test.description, test.isbn, test.expected, observed)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func IsValidISBN(s string) bool {
	var nums []int
	n := len(s)
	for i := 0; i < n; i++ {
		if i == n - 1 && s[i] == 'X'{
			nums = append(nums, 10)
			break
		}
		if '0' <= s[i] && s[i] <= '9'{
			nums = append(nums, int(s[i] - '0'))
		} else if s[i] != '-' {
			return false
		}
	}

	if len(nums) != 10 {
		return false
	}
	sum := 0
	for i,j := 10,0; i > 0; i-- {
		sum += nums[j] * i
		j++
	}
	return sum % 11 == 0
}

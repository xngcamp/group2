package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Solution("((1)23(45))(aB)", 0))
}

func Solution(str string, idx int) (int, error) {
	var stack []byte
	var end int
	status := false
	for i, c := range str {
		if i == idx {
			if c == '(' {
				status = true
			} else {
				return -1, errors.New("Not an opening bracket")
			}
		}
		if status {
			if c == '(' {
				stack = append(stack, '(')
			} else if c == ')' {
				n := len(stack)
				if n <= 1 {
					stack = nil
				} else {
					stack = stack[:n-1]
				}
			}
			if len(stack) == 0 {
				end = i
				break
			}
		}
	}

	if status {
		if end != 0 {
			return end, nil
		}
		return 0, errors.New("Not an closeing bracket")
	}

	return 0, errors.New("Not an opening bracket")
}



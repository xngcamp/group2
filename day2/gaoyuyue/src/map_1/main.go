package main

import "fmt"

func main() {
	str := CalMaxUniqueStr("abcdefbgac")
	fmt.Println(str)
}

func CalMaxUniqueStr(str string) string {
	set := make(map[byte]int)
	begin := 0
	var result string

	n := len(str)
	for i := 0; i < n; i++ {
		if v, ok := set[str[i]]; ok {
			if i - begin > len(result) {
				result = str[begin:i]
			}
			deleteSome(set, str[begin:v + 1])
			begin = v + 1
		}
		set[str[i]] = i
	}

	if n - begin > len(result){
		result = str[begin:n]
	}

	return result
}

func deleteSome(set map[byte]int, s string) {
	for i := 0; i < len(s); i++ {
		delete(set, s[i])
	}
}

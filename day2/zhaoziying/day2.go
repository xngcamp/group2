package awesomeProject2

import "fmt"

func lengthOfNonRepeatingSubstr(s tring) int{
	lastOccurred := make(map[rune]int)
	start :=0
	maxLength :=0

	for i, ch := range []rune(s){
		if lastI,ok := lastOccurred[ch]; ok && lastI >= start{
			start = lastI + 1
		}
		if i-start + 1 >maxlength{
			maxlength = i -start + 1
		}
		lastOccurred[ch]=i
	}
	return  maxlength
}

func main( ){
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("  "))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdef"))

}


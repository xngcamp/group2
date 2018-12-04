package map1

import "testing"

func TestGetMaxString(t *testing.T) {
	var tests = []struct{
		_str string
		_outstr string
	}{
		{"abcdefabc","abcdef"},
		{"abcdefabcdefgh","abcdefgh"},
		{"ababab","ab"},
	}

	for _,test := range tests{
		_result := GetMaxString(test._str)
		if test._outstr != _result{
			t.Error("the outstring is wrong")
		}
	}
}
package slice1

import "testing"

func TestGetIndex(t *testing.T) {

	var tests = []struct{
		_str string
		_index int
		_result int
	}{
		{"((a)(b)(c))",0,10},
		{"((a)(b)(c))",1,3},
		{"((a)(b)(c))",4,6},
	}

	for _,test := range tests{
		_res := GetIndex(test._str,test._index)
		if _res != test._result {
			t.Error("the string index is wrong")
		}
	}
}

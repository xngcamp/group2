package slice1

func GetIndex(_str string, _index int) int {
	_result := 0
	_count := 0
	for i:=_index;i< len(_str);i++  {
		//fmt.Println(string(_str[i]))
		if string(_str[i]) == "(" {
			_count++
		}else if string(_str[i]) == ")" {
			_count--
		}
		if _count == 0 {
			_result = i
			break
		}
	}
	return _result
}

package map1

func GetMaxString(str string) string {
	_currstr := ""
	_map := map[string]int{}
	_length :=0
	_key := ""
	for i:=0;i< len(str);i++  {
		for j:=i;j< len(str);j++  {
			if int(str[i]) - i == int(str[j]) -j{
				_currstr += string(str[j])
			}
		}
		_map[_currstr] = len(_currstr)
		_currstr = ""
	}

	for key,value := range _map{
		if _length < value {
			_length = value
			_key = key
		}
	}
	return _key
}


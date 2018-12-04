package tmp

import (
	"errors"
)

func Solution(str string, idx int) (int, error) {
	b := []rune(str) //将字符串转换成字节数组
	count := 0
	//如果字符串指定索引处不是左括号，返回-1
	if b[idx] != 40 {
		return -1, errors.New("Not an opening bracket")
	}
	//从左括号开始遍历，寻找第一个匹配的右括号
	//思路：从idx+1处开始遍历，如果第一个出现左括号，那么跳过随之出现的右括号
	//		如果第一个出现右括号，则返回此处索引
	//		count表示需要跳过的右括号个数
	for i := idx + 1; i < len(b); i++ {
		//如果idx+1处后出现的第一个符号是左括号，那么代表出现嵌套括号
		//需要跳过的右括号+1
		if b[i] == 40 {
			count++
		}
		//idx后出现匹配右括号
		if count == 0 && b[i] == 41 {
			return i, errors.New("Found an closing bracket")
		}
		//出现一个右括号，与idx索引后面出现的第一个左括号匹配
		//因此抵消，count--
		if b[i] == 41 {
			count--
		}
	}
	//没有结果
	return 0, errors.New("Not an opening bracket")
}

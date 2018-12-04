package isbn

import (
	"strings"
)

/**
思路：
	1、首先将字符串按照“-”分隔开
	逐个判断分割得到的子字符串，如果存在除“-”、“X”、0-9数字之外的字符，返回false
	否则，将子字符串追加到临时定义的slice中
	2、然后，判断临时定义的slice与分割得到的字符串数组长度是否相同，相同代表isbn字符串不包含异常字符
	3、将临时定义的slice转换为字符串，判断字符串长度是否符合isbn规则
	4、利用mod公式判断isbn是否符合规则
	5、返回结果
*/

func IsValidISBN(isbn string) bool {
	flag := false                 //标志变量
	b := strings.Split(isbn, "-") //将isbn字符串分割
	var result []string           //定义临时slice，保存分割的子字符串
	var temp string               //保存临时slice转换成的字符串
	for index := range b {        //判断子字符串是否包含异常字符
		byteArr := []rune(b[index])
		i := 0
		for ; i < len(byteArr); i++ {
			if !((byteArr[i] >= 48 && byteArr[i] <= 57) || byteArr[i] == 88) {
				break
			}
		}
		if i == len(byteArr) {
			result = append(result, b[index])
		} else {
			break
		}
	}
	if len(b) == len(result) { //不存在异常字符，将字符数组转换为字符串，否则直接返回false
		temp = strings.Join(result, "")
	}
	if strings.Count(temp, "")-1 == 10 { //字符串长度为10，符合要求，否则直接返回false
		sum := 0
		tempByte := []byte(temp)
		for index, value := range tempByte { //开始计算mod值
			if index != len(tempByte)-1 && value == 88 { //如果X不是在最后
				flag = false
				return flag
			} else if index == len(tempByte)-1 && value == 88 { //如果X在最后
				sum = sum + 10*(len(temp)-index)
			} else { //其他
				sum = sum + int((value-'0'))*(len(tempByte)-index)
			}
		}
		if sum%11 == 0 { //如果mod11等于0，表示符合ibsn规则，返回true
			flag = true
		}
	}
	return flag
}

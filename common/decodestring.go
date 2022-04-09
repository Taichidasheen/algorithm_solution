package common

import "strconv"

/*
https://leetcode-cn.com/problems/decode-string/
 */

func decodeString(s string) string {
	res, _ := _decodeString([]byte(s))
	return string(res)
}

func _decodeString(bytes []byte) ([]byte, int) {
	var res []byte
	var multi int
	var cnt int
	for idx:=0;idx<len(bytes);idx++ {
		char := bytes[idx]
		if char >='a' && char <='z' {
			res = append(res, char)
			cnt ++
		}
		if char >= '0' && char <= '9' {
			num,_ := strconv.Atoi(string(char))
			multi = multi * 10 + num
			cnt ++
		}
		if char == '[' {
			newStr := bytes[idx+1:]
			recur, recnt := _decodeString(newStr)
			for i:=0;i<multi;i++ {
				res = append(res, recur...)
			}
			cnt ++
			cnt = cnt + recnt
			idx = idx + recnt
			multi = 0//这里一定要记得将multi重新置为0，否则对于输入3[a]2[bc]，会出现multi=32的情况
		}
		if char == ']' {
			cnt ++
			return res, cnt
		}
	}
	return res, cnt
}
package recursive

/*
剑指 Offer II 086. 分割回文子字符串
https://leetcode.cn/problems/M99OJA/
解法：
递归+备忘录
*/

func partition(s string) [][]string {
	res := make(map[string][][]string)
	_partition(s, res)
	return res[s]
}

func _partition(s string, res map[string][][]string) {
	if len(s) == 1 {
		res[s] = [][]string{[]string{s}}
		return
	}
	if isPalindrome(s) {
		res[s] = append(res[s], []string{s})
	}
	for i:=1;i<len(s);i++ {
		head := s[0:i]
		tail := s[i:]
		if isPalindrome(head) {
			if res[tail] != nil {
				input := res[tail]
				temp := appendFront(head, input)
				res[s] = append(res[s], temp...)
				//fmt.Printf("res[%s]:%+v\n", s, res[s])
			} else {
				_partition(tail, res)
				input := res[tail]
				temp := appendFront(head, input)
				res[s] = append(res[s], temp...)
				//fmt.Printf("res[%s]:%+v\n", s, res[s])
			}
		}
	}
}

func isPalindrome(s string) bool {
	//fmt.Printf("isPalindrome:%s\n", s)
	i := 0
	j := len(s)-1
	for i<=j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	//fmt.Printf("isPalindrome:%v\n", true)
	return true
}

func appendFront(head string, input [][]string) [][]string {
	var res [][]string
	for _, row := range input {
		var temp []string
		temp = append(temp, head)
		temp = append(temp, row...)
		res = append(res, temp)
	}
	return res
}

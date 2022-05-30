package recursive

/*
剑指 Offer II 085. 生成匹配的括号
https://leetcode.cn/problems/IDBivT/
解法：
递归+排列
感悟：只有当当前面临多个选择的时候，才需要考虑路径回退，如果只有一个选择，则不需要考虑路径回退
比如下面的解法中，只有left<right时，可以选择左括号或右括号，其他条件下只能做一种选择
*/

func generateParenthesis(n int) []string {
	var path []uint8
	var res []string
	_generateParenthesis(n, n, path, &res)
	return res
}

func _generateParenthesis(left, right int, path []uint8, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, string(path))
		return
	}
	if left == 0 && right > 0 {
		//只能选right
		path = append(path, ')')
		_generateParenthesis(left, right-1, path, res)
		return
	}
	if left < right {
		//选择left
		var tempPath1 []uint8
		tempPath1 = append(tempPath1, path...)
		tempPath1 = append(tempPath1, '(')
		_generateParenthesis(left-1, right, tempPath1, res)

		//选择right
		var tempPath2 []uint8
		tempPath2 = append(tempPath2, path...)
		tempPath2 = append(tempPath2, ')')
		_generateParenthesis(left, right-1, tempPath2, res)
		return
	}
	if left == right {
		//只能选left
		path = append(path, '(')
		_generateParenthesis(left-1, right, path, res)
		return
	}
}

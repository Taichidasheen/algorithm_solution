package recursive

/*
剑指 Offer II 083. 没有重复元素集合的全排列
https://leetcode.cn/problems/VvJkup/
解法：
递归+排列
解法的重要点：
used map[int]bool，int是nums数组元素的下标，而不是nums数组元素的值
*/

func permute(nums []int) [][]int {
	var path []int
	used := make(map[int]bool)
	var res [][]int
	_permute(nums, path, used, &res)
	return res
}

func _permute(nums []int, path []int, used map[int]bool, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i:=0;i<len(nums);i++ {
		if used[i] == false {
			path = append(path, nums[i])
			used[i] = true
			_permute(nums, path, used, res)
			used[i] = false
			path = path[0:len(path)-1]
		}
	}
}
package recursive

import "sort"

/*
剑指 Offer II 084. 含有重复元素集合的全排列
https://leetcode.cn/problems/7p8L0Z/
解法：
递归+排列
解法的重要点：
used map[int]bool，int是nums数组元素的下标，而不是nums数组元素的值
遇到数组中相同的元素的处理方法：如果当前元素的值不唯一，则看与当前元素值相同的上一个元素是否被使用了，
如果被使用了，则可以选择当前元素，如果没有被使用，则不能选择当前元素。（前提是按照数组中元素的顺序依次进行选择）
*/


func permuteUnique(nums []int) [][]int {
	var res [][]int
	used := make(map[int]bool)
	var path []int
	sort.Ints(nums)
	_permuteUnique(nums, path, used, &res)
	return res
}

func _permuteUnique(nums []int, path []int, used map[int]bool, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i:=0;i<len(nums);i++ {
		if  i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		if used[i] == false {
			path = append(path, nums[i])
			used[i] = true
			_permuteUnique(nums, path, used, res)
			used[i] = false
			path = path[0:len(path)-1]
		}
	}
}
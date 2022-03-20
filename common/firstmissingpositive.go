package common

/*
https://leetcode-cn.com/problems/first-missing-positive/

 */

func firstMissingPositive(nums []int) int {
	//第一次遍历把负数都置成0
	for i, num := range nums {
		if num <= 0 {
			nums[i] = 1000000
		}
	}
	//第二次遍历把nums[i]对应的nums数组位置的数置成相反数
	for _, num := range nums {
		absNum := abs(num)
		//注意：这里一定加上nums[absNum - 1] > 0条件，不然对于输入{1，1}的case通过不了
		if absNum > 0  && absNum <= len(nums) && nums[absNum - 1] > 0 {
			nums[absNum - 1] = -1 * nums[absNum - 1]
		}
	}
	//第三次遍历返回第一个nums[i] > 0的下标i
	for i, num := range nums {
		if num >= 0 {
			return i+1
		}
	}
	return len(nums) + 1
}

func abs(v int) int {
	if v >= 0 {
		return v
	}
	return v * -1
}
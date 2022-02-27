package dp

import "fmt"

/*
lc-53: https://leetcode.com/problems/maximum-subarray/
 */

func maxSubArray(nums []int) int {

	return dp(nums)
	//return slidingWindow(nums)
	//return bruteForce(nums)

}

//暴力解法
func bruteForce(nums []int) int {

	len := len(nums)
	var total [][]int
	for i:=0;i<len;i++{
		var row []int
		for j:=0;j<len;j++ {
			row = append(row, 0)
		}
		total = append(total, row)
	}

	for i:=0;i<len;i++ {
		total[i][i] = nums[i]
		for j:=i+1;j<len;j++ {
			total[i][j] = total[i][j-1] + nums[j]
		}
	}
	var maxValue int
	for _, row := range total {
		for _, val := range row {
			if val > maxValue {
				maxValue = val
			}
		}
	}
	return maxValue
}

//滑动窗口
func slidingWindow(nums []int) int {
	var sum int
	maxSum := nums[0]
	var left, right int
	for right < len(nums) {
		sum = sum + nums[right]
		maxSum = max(maxSum, sum)
		//扩张窗口
		if sum > 0 {
			right ++
		} else {
			//缩减窗口
			for sum <= 0 && left < right {
				sum = sum - nums[left]
				left ++
			}
			//加上下面这个逻辑是因为如果数组中有大于0的数，那最大子数组不会以负数开头
			if left == right && nums[left] <= 0 {
				left ++
				right ++
			}
			sum = 0

		}
	}
	return maxSum
}

//动态规划
func dp(nums []int) int {
	len := len(nums)
	var dp []int = make([]int, len)
	dp[0] = nums[0]
	for i:=1;i<len;i++ {
		//以nums[i]结尾的最大子数组dp[i]要么为nums[i]要么为dp[i-1]+nums[i]
		dp[i] = max(nums[i], dp[i-1] + nums[i])
		fmt.Printf("dp[%d]:%d\n",i, dp[i])
	}

	maxValue := dp[0]
	for _, val := range dp {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}

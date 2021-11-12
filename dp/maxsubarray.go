package dp

/*
lc-53: https://leetcode.com/problems/maximum-subarray/
 */

func maxSubArray(nums []int) int {

	return dp(nums)
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
	return 0
}

//动态规划
func dp(nums []int) int {
	len := len(nums)
	var dp []int = make([]int, len)
	dp[0] = nums[0]
	for i:=1;i<len;i++ {
		dp[i] = max(nums[i], dp[i-1] + nums[i])
	}

	maxValue := dp[0]
	for _, val := range dp {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}

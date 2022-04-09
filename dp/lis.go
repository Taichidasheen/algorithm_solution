package dp

/*
 lc-300: https://leetcode.com/problems/longest-increasing-subsequence/
dp[i]的定义：以nums[i]为结尾的最长递增子序列长度
 */

//使用迭代方法
func lengthOfLIS(nums []int) int {
	len := len(nums)
	var dp []int
	for i := 0; i < len; i++ {
		dp = append(dp, 1)
	}

	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	var maxLen int
	for i := 0; i < len; i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}

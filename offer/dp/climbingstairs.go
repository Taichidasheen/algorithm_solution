package dp

/*
剑指 Offer II 088. 爬楼梯的最少成本
https://leetcode.cn/problems/GzCJIP/
解法：
dp
dp[i]表示到达第i级台阶时的最小成本
dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
*/

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i:=2;i<len(cost);i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1],dp[len(cost)-2])
}
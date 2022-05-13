package dp

/*
剑指 Offer II 091. 粉刷房子
https://leetcode.cn/problems/JEj789/
解法：
dp,颜色顺序为红，蓝，绿
定义dp[i][0]为将第i个房子粉刷成红色的最小成本
定义dp[i][1]为将第i个房子粉刷成蓝色的最小成本
定义dp[i][2]为将第i个房子粉刷成绿色的最小成本
*/

func minCost(costs [][]int) int {
	num := len(costs)
	dp := make([][]int, num)
	for i:=0;i<num;i++ {
		dp[i] = make([]int, 3)
	}
	dp[0] = costs[0]
	for i:=1;i<len(costs);i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}
	return min(min(dp[num-1][0], dp[num-1][1]), dp[num-1][2])
}

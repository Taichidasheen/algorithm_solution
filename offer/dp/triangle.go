package dp

/*
剑指 Offer II 100. 三角形中最小路径之和
https://leetcode.cn/problems/IlPe0q/
解法：
dp[i][j] = min(dp[i-1][j],dp[i-1][j-1])
https://leetcode-cn.com/problems/triangle/solution/san-jiao-xing-zui-xiao-lu-jing-he-by-leetcode-solu/
 */

func minimumTotal(triangle [][]int) int {
	var dp [][]int
	for i:=0;i<len(triangle);i++ {
		row := make([]int, len(triangle[i]))
		dp = append(dp, row)
	}

	dp[0][0] = triangle[0][0]
	for i:=1;i<len(triangle);i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
	}
	for i:=1;i<len(triangle);i++ {
		for j:=1;j<len(triangle[i]);j++ {
			if j == len(triangle[i]) - 1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	minSum := dp[len(triangle)-1][0]
	for _, val := range dp[len(triangle)-1] {
		minSum = min(minSum, val)
	}
	return minSum
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
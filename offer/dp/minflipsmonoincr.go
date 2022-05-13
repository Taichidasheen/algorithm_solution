package dp

/*
剑指 Offer II 092. 翻转字符
https://leetcode.cn/problems/cyJERH/
解法：
dp，参考评论区，官方解答不够清晰。dp[i]定义为以s[i]结尾的字符串的最小的翻转次数
https://leetcode.cn/problems/flip-string-to-monotone-increasing/solution/jiang-zi-fu-chuan-fan-zhuan-dao-dan-diao-di-zeng-b/
*/

func minFlipsMonoIncr(s string) int {
	dp := make([]int, len(s))
	var oneCnt int
	dp[0] = 0//无论第一个字符是0还是1，都不需要翻转
	if s[0] == '1' {
		oneCnt++
	}
	for i:=1;i<len(s);i++ {
		if s[i]== '1' {
			oneCnt++
			dp[i] = dp[i-1]
		} else {
			dp[i] = min(dp[i-1]+1, oneCnt)
		}
	}
	return dp[len(s)-1]
}


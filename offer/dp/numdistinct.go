package dp

/*
剑指 Offer II 097. 子序列的数目
https://leetcode.cn/problems/21dk04/
解法：
dp，注意：官方题解是从前往后匹配的
https://leetcode-cn.com/problems/distinct-subsequences/solution/bu-tong-de-zi-xu-lie-by-leetcode-solutio-urw3/
其实也可以是从后往前匹配，转移方程为：
dp[i][j] = dp[i-1][j-1] + dp[i-1][j] if s[i]=s[j]，s是源字符串，t是目标字符串
dp[i][j] = dp[i-1][j] if s[i]!=s[j]
注意边界：
if i<j dp[i][j]=0, 无法通过s转化成t
dp[i][0] = dp[i-1][0] if s[i] != t[0]
dp[i][0] = dp[i-1][0] + 1 if s[i] = t[0]
dp[0][0] = 0, if s[i] != t[0]
dp[0][0] = 1, if s[i] == t[0]
s和t还有可能是空字符串，如果s不为空，t为空，结果等于1；如果s为空，t不为空，结果等于0；如果都为空，结果等于1
 */

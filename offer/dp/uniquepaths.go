package dp


/*
剑指 Offer II 098. 路径的数目
https://leetcode.cn/problems/2AoeFn/
解法：
1. dp.
dp[i][j] = dp[i-1][j] + dp[i][j-1]
2. 排列组合计算
A(m-1+n-1)中取m-1+n-1的全排列/（A m-1中取m-1的排列*An-1中取n-1的排列）
https://leetcode-cn.com/problems/unique-paths/solution/bu-tong-lu-jing-by-leetcode-solution-hzjf/
 */

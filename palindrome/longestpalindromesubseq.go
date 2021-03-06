package palindrome

/*
lc-516: https://leetcode.com/problems/longest-palindromic-subsequence/
 */
func longestPalindromeSubseq(s string) int {
	//反转字符串
	var rev  []byte
	for i:= len(s)-1;i>=0;i-- {
		rev = append(rev, s[i])
	}
	revstr := string(rev)
	return longestCommonSubsequence(s, revstr)
}

//使用迭代方法
func longestCommonSubsequence(text1 string, text2 string) int {
	//初始化数组
	len1 := len(text1)
	len2 := len(text2)
	var dp [][]int
	for i := 0; i <= len1; i++ {
		var temp []int
		for j := 0; j <= len2; j++ {
			temp = append(temp, 0)
		}
		dp = append(dp, temp)
	}
	//fmt.Printf("dp:%+v\n", dp)

	for i := 1; i <=len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len1][len2]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
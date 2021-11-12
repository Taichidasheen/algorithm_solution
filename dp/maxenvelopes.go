package dp

import "sort"


/*
lc-354:https://leetcode.com/problems/russian-doll-envelopes/submissions/
 */
func maxEnvelopes(envelopes [][]int) int {
	//初始化dp数组
	len := len(envelopes)
	var dp []int
	for i := 0; i < len; i++ {
		dp = append(dp, 1)
	}

	sorted := sortEnvelops(envelopes)

	for i := 0; i < len; i++ {
		for j := 0; j < i; j++ {
			if sorted[i][1] > sorted[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	var maxNum int
	for i := 0; i < len; i++ {
		if dp[i] > maxNum {
			maxNum = dp[i]
		}
	}
	return maxNum

}

//比较信封大小
func isEnvelopAFitB(a, b []int) bool {
	if a[0] > b[0] && a[1] > b[1] {
		return true
	}
	return false
}

//对信封进行排序
func sortEnvelops(envelops [][]int) [][]int {
	sort.Slice(envelops[:], func(i, j int) bool {
		if envelops[i][0] > envelops[j][0] {
			return false
		} else if envelops[i][0] == envelops[j][0] {
			return envelops[i][1] > envelops[j][1]
		} else {
			return true
		}
	})
	return envelops
}
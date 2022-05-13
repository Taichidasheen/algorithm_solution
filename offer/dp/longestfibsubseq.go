package dp

/*
剑指 Offer II 093. 最长斐波那契数列
https://leetcode.cn/problems/Q91FMA/
解法：
直接暴力解法比较靠谱，dp不容易理解，时间复杂度也差不多
https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence/solution/zui-chang-de-fei-bo-na-qi-zi-xu-lie-de-chang-du-by/
*/

func lenLongestFibSubseq(arr []int) int {
	arrMap := make(map[int]int)
	for _, val := range arr {
		arrMap[val] = 1
	}
	var maxLength int
	for i:=0;i<len(arr);i++ {
		for j:=i+1;j<len(arr);j++ {
			pre := arr[i]
			cur := arr[j]
			sum := pre + cur
			length := 2
			for arrMap[sum] > 0{
				length++
				pre = cur
				cur = sum
				sum = pre + cur
			}
			maxLength = max(maxLength, length)
		}
	}
	if maxLength < 3 {
		return 0
	}
	return maxLength
}
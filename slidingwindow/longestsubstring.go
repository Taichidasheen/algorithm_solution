package slidingwindow

/*
 无重复字符最长子串
 lc-3: https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/
func lengthOfLongestSubstring(s string) int {
	unique := make(map[uint8]int)
	var left, right int
	var start, end, max int
	for right < len(s) {
		ch := s[right]
		if unique[ch] == 0 {
			//没有遇到重复情况,滑动窗口扩大
			right++
			unique[ch]++
			//注意：这里最大长度的计算要放到这里，不能放到else结构体里，不然对于输入s="x"这样的case会有问题
			//计算最大长度
			end = right - 1
			start = left
			//下面这段可以简单写成 max = max(max, end-start+1)
			if end-start + 1 > max {
				max = end-start + 1
			}
		} else {
			//遇到重复情况，滑动窗口收缩
			for s[left] != s[right] {
				unique[s[left]] = 0
				left++
			}
			unique[s[left]] = 0
			left++
		}
	}
	return max
}

package practice

//使用迭代法
func iterReverseList(node *ListNode) *ListNode {
	var prev, cur *ListNode
	cur = node
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

//使用递归法
func recursionReverseList(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}
	head := recursionReverseList(node.Next)
	node.Next.Next = node
	node.Next = nil
	return head
}

//最长增长子序列
//定义dp[i]为以nums[i]结尾的最长子序列长度
//动态规划方程：
// dp[0] = 1
// dp[i] = max(dp[n] + 1, dp[i]) 如果nums[i] > nums[n] (0<=n<i-1)
func lis(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = 1
	for i:= 1;i<length;i++ {
		for j:=0;j<i;j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	var maxLength int
	for _, val := range dp {
		maxLength = max(maxLength, val)
	}
	return maxLength
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 无重复字符最长子串
思路:使用滑动窗口
 */
func lengthOfLongestSubstring(s string) int {
	bytes := []byte(s)
	dup := make(map[byte]int)
	var left, right int
	var maxLength int
	for right < len(bytes) {
		//增长窗口
		if dup[bytes[right]] == 0 {
			right++
			//注意这个地方，right要减去1再计算长度
			maxLength = max(maxLength, right -1 -left + 1)
			dup[bytes[right]]++
			continue
		}
		//缩减窗口，需要用循环来做，缩减到没有重复元素
		for dup[bytes[right]] != 0 {
			dup[bytes[left]]--
			left ++
		}
	}
	return maxLength
}

package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
https://leetcode.com/problems/swap-nodes-in-pairs/
思路1： 可以使用n个一组反转链表的思想，对于这道题就是n=2
思路2： 使用递归的思想，每次交换链表的头两个节点。
 */

//思路1解法
func swapPairs_1(head *ListNode) *ListNode {
	return nil
}

//迭代法反转链表前N个节点
func reverseTopN(head *ListNode, n int) *ListNode {
	var prev, cur *ListNode
	cur = head
	for i:=0;i<n;i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	head.Next= cur
	return prev
}











//思路2解法
func swapPairs_2(head *ListNode) *ListNode {
	return nil
}

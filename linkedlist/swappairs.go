package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
https://leetcode.com/problems/swap-nodes-in-pairs/
思路1： 可以使用n个一组反转链表的思想，对于这道题就是n=2
参考：https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/
思路2： 使用递归的思想，每次交换链表的头两个节点。
 */

//思路1解法
func swapPairs_1(head *ListNode) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < 2; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next //临时保存next节点
		head, tail = reverseHeadTail(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

//反转head到tail区间的链表
func reverseHeadTail(head, tail *ListNode) (*ListNode, *ListNode) {
	tailNext := tail.Next
	var prev, cur *ListNode
	cur = head
	for cur != tailNext {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return tail, head
}


//思路2解法
func swapPairs_2(head *ListNode) *ListNode {
	return nil
}

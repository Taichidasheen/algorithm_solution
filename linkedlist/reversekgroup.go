package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	//检查链表长度是否大于等于K
	var tail *ListNode
	tail = head
	for i:=1;i< k; i++ {
		if head.Next == nil {
			return head
		}
		tail = tail.Next
	}
	successor = tail.Next
	newHead := reverseBetween(head, tail)
	head.Next = reverseKGroup(successor, k)
	return newHead
}

//反转head和tail之间的元素。注意：迭代方法反转链表时for结构里面只有4句，只涉及一次指针转移操作
func reverseBetween(head *ListNode, tail *ListNode) *ListNode {
	var pre, cur *ListNode
	cur = head
	for cur != tail.Next {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
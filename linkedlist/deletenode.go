package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
思路：双指针
1. 先让一个指针走n步，然后另一个指针开始走
2. 先走的指针到达末尾的时候（指向nil），后走的指针指向倒数第n个节点
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var ptr1, ptr2 *ListNode
	ptr1 = head
	ptr2 = head
	for i:=0;i<n;i++ {
		if ptr1 == nil {
			//还没有走n步，就已经抵达链表末尾，直接return head
			return head
		}
		ptr1 = ptr1.Next
	}
	var beforePtr2 *ListNode
	beforePtr2 = head
	//下面3行代码是为了处理链表长度为n,删除倒数第n个节点的情况（也就是说删除头节点）
	if ptr1 == nil {
		return head.Next
	}
	//ptr1，ptr2先各走一步
	ptr1 = ptr1.Next
	ptr2 = ptr2.Next
	for ptr1 != nil {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
		beforePtr2 = beforePtr2.Next
	}
	beforePtr2.Next = ptr2.Next
	return head
}
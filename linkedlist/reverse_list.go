package linkedlist


//reverseLinkedList 反转链表 迭代方式
func reverseLinkedList(node *ListNode) *ListNode {
	pre := node
	next := node.Next
	pre.Next = nil
	for next != nil {
		nextnext := next.Next
		next.Next = pre
		pre = next
		next = nextnext
	}
	return pre
}

//recursiveReverseLinkedList 反转链表 递归方式
func recursiveReverseLinkedList(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}
	head := recursiveReverseLinkedList(node.Next)
	node.Next.Next = node
	node.Next = nil
	return head
}


var successor *ListNode
//reverseListRecursionTopN 递归方式反转链表前N个节点
func reverseListRecursionTopN(node *ListNode, topN int64) *ListNode {
	if topN == 1 {
		successor = node.Next
		return node
	}
	//head指向反转链表的头节点
	head := reverseListRecursionTopN(node.Next, topN-1)
	node.Next.Next = node
	node.Next = successor
	return head
}

//reverseListRecursionBetween 递归方式反转链表a-b区间的节点
func reverseListRecursionBetween(node *ListNode, a int64, b int64) *ListNode {

	if a == 1 {
		return reverseListRecursionTopN(node, b)
	}
	halfTail := node
	i := a
	for i > 2 {
		halfTail = node.Next
		i = i - 1
	}
	//head指向反转链表的头节点
	head := reverseListRecursionTopN(halfTail.Next, b-a+1)
	halfTail.Next = head
	return node
}
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

//reverseLinkedListTopN 反转链表前N个节点
func reverseLinkedListTopN(node *ListNode, topN int32) *ListNode {
	return nil

}


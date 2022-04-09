package practice

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println()
}

//UnmarshalListBySlice 根据数组反序列化链表
func UnmarshalListBySlice(nums []int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v, Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}

//迭代反转1
func reverseList1(node *ListNode) *ListNode {
	var pre, cur *ListNode
	cur = node
	for cur.Next != nil {
		//1. 先临时获取cur的下一个节点next
		next := cur.Next
		//2. 将当前节点的指针指向pre节点
		cur.Next = pre
		//3. 更新cur节点和pre节点
		pre = cur
		cur = next
	}
	//注意：最后一个不为空的节点需要做一次指针变换操作，从指向nil改为指向pre,最后返回cur
	cur.Next = pre
	return cur
}

//迭代反转2
func reverseList2(node *ListNode) *ListNode {
	var pre, cur *ListNode
	cur = node
	for cur != nil {
		//1. 先临时获取cur的下一个节点next
		next := cur.Next
		//2. 将当前节点的指针指向pre节点
		cur.Next = pre
		//3. 更新cur节点和pre节点
		pre = cur
		cur = next
	}
	//注意：和reverseList1不同的地方是最后返回的是pre，因为返回之前pre指向最后一个不为空的节点，而cur指向最后一个nil
	return pre
}

//递归反转
func recursiveList(node *ListNode) *ListNode {
	//下面这个if判断不应该要，应该判断node.next
	/*if node == nil {
		return node
	}*/
	if node.Next == nil {
		return node
	}
	//1. 先进行递归
	head := recursiveList(node.Next)
	//2. 更新节点的指向
	node.Next.Next = node
	node.Next = nil
	return head
}

//迭代反转前n个节点
func reverseTopN(node *ListNode, n int) *ListNode {
	if n <= 0 {
		return node
	}
	var pre, cur *ListNode
	cur = node
	for i:=1; i<=n; i++ {
		//1.保存当前节点的下一个节点
		next := cur.Next
		//2. 更改当前节点的next指针
		cur.Next = pre
		//3. 更新pre,cur指针指向
		pre = cur
		cur = next
	}
	//将已经反转的部分连接到未反转的部分上
	node.Next = cur
	return pre
}

//递归反转前n个节点(比迭代法反转要难于理解一些)
func recursiveTopN(node *ListNode, n int) *ListNode {
	if node.Next == nil {
		return node
	}
	if n <= 1 {
		return node
	}
	head := recursiveTopN(node.Next, n-1)
	//要想象经过这一步之后，node.Next后面的节点已经是反转好的（node.Next不一定是头节点）
	//接下来要做的就是断开node.Next和它的后继节点successor的指针，
	//将node.Next指向当前node， 并将当前节点node指向后继节点successor

	//1. 保存node.Next节点指向的下一个节点
	successor := node.Next.Next
	//2. 让node.Next节点指向node节点
	node.Next.Next = node
	//3. 让node节点指向successor节点
	node.Next = successor
	return head
}

//反转区间内的
func reverseBetween(node *ListNode, a int, b int) *ListNode {
	if b < a {
		return node
	}
	if a <= 1 {
		return reverseTopN(node, b)
	}
	halfTail := node
	for i:=2;i<a;i++ {
		halfTail = halfTail.Next
	}
	head := reverseTopN(halfTail.Next, b-a + 1)
	halfTail.Next = head
	return node
}



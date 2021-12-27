package linkedlist

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

func UnmarshalListBySlice2(nums []int) *ListNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	if length >= 2 {
		cur := head
		for i := 1; i <= length-1; i++ {
			cur.Next = &ListNode{nums[i], nil}
			cur = cur.Next
		}
	}
	return head
}
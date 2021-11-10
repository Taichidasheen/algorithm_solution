package structure

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//IsEqual	比较结果
func IsEqual(l1 *ListNode, l2 *ListNode) bool {

	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 != nil {
		return false
	}
	if l1 != nil && l2 == nil {
		return false
	}
	return true
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

// UnmarshalListByRand 随机初始化链表
func UnmarshalListByRand(max_num int, len int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head

	for i := 0; i < len; i++ {
		tmp.Next = &ListNode{Val: rand.Intn(max_num), Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}

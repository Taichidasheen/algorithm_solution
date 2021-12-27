package linkedlist

import (
	"testing"
)

func Test_reverseLinkedList(t *testing.T) {

	arr := []int{5,4,3,2,1}
	list := UnmarshalListBySlice(arr)
	PrintList(list)

	reverseList := reverseLinkedList(list)
	PrintList(reverseList)
}

func Test_recursiveReverseLinkedList(t *testing.T) {
	arr := []int{5,4,3,2,1}
	list := UnmarshalListBySlice(arr)
	PrintList(list)

	reverseList2 := recursiveReverseLinkedList(list)
	PrintList(reverseList2)
}
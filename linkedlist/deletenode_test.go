package linkedlist

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	//arr := []int{5,4,3,2,1}
	arr := []int{1}
	list := UnmarshalListBySlice(arr)
	PrintList(list)
	list = removeNthFromEnd(list, 1)
	PrintList(list)
}

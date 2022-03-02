package practice

import (
	"testing"
)

//迭代方法反转
func Test_reverseList(t *testing.T) {
	arr := []int{5,4,3,2,1}
	list := UnmarshalListBySlice(arr)
	PrintList(list)
	//reverse := reverseList1(list)
	reverse := reverseList2(list)
	PrintList(reverse)
}

//递归方法反转
func Test_recursiveList(t *testing.T) {
	arr := []int{5,4,3,2,1}
	list := UnmarshalListBySlice(arr)
	PrintList(list)
	reverse := recursiveList(list)
	PrintList(reverse)
}

//迭代方法反转
func Test_reverseTopN(t *testing.T) {
	arr := []int{5,4,3,2,1}
	list := UnmarshalListBySlice(arr)
	PrintList(list)
	reverse := reverseTopN(list, 4)
	PrintList(reverse)
}

//递归方法反转
func Test_recursiveTopN(t *testing.T) {
	arr := []int{5,4,3,2,1}
	list := UnmarshalListBySlice(arr)
	PrintList(list)
	reverse := recursiveTopN(list, 4)
	PrintList(reverse)
}

func Test_reverseBetween(t *testing.T) {
	arr := []int{5,4,3,2,1}
	list := UnmarshalListBySlice(arr)
	PrintList(list)
	reverse := reverseBetween(list, 2, 4)
	PrintList(reverse)
}
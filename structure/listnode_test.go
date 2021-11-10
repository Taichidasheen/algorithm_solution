package structure

import "testing"

func TestUnmarshalListByRand(t *testing.T) {

}

func TestUnmarshalListBySlice2(t *testing.T) {
	t.Logf("run test...")
	arr := []int{10}
	list1 := UnmarshalListBySlice(arr)
	t.Logf("UnmarshalListBySlice list1:")
	PrintList(list1)
	list2 := UnmarshalListBySlice2(arr)
	t.Logf("UnmarshalListBySlice2 list2:")
	PrintList(list2)
}

package sort

import (
	"testing"
)

func Test_bubbleSort(t *testing.T) {

	arr := []int{5,4,3,2,1,8,9}
	t.Logf("input arr:%+v", arr)
	res := bubbleSort(arr)
	t.Logf("res:%+v", res)
}

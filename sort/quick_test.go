package sort

import (
	"testing"
)

func Test_quickSort(t *testing.T) {
	//arr := []int{5,4,3,2,1,8,9,4,10,2}
	arr := []int{10,9,8,7,6,5}
	t.Logf("input arr:%+v", arr)
	res := quickSort(arr)
	t.Logf("res:%+v", res)
}

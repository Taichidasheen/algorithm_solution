package dp

import "testing"

func Test_lenLongestFibSubseq(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8}
	t.Logf("%d", lenLongestFibSubseq(arr))

	arr = []int{1,3,7,11,12,14,18}
	t.Logf("%d", lenLongestFibSubseq(arr))

	arr = []int{2,4,7,8,9,10,14,15,18,23,32,50}
	t.Logf("%d", lenLongestFibSubseq(arr))
}
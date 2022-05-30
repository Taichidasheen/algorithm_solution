package recursive

import (
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	nums := []int{3,3,0,3}
	res := permuteUnique(nums)
	t.Logf("res:%+v", res)
}
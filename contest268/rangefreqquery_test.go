package contest268

import "testing"

func TestRangeFreqQuery_Query(t *testing.T) {
	arr := []int{12,33,4,56,22,2,34,33,22,12,34,56}
	con := Constructor(arr)
	t.Logf("con:%+v", con)
	res := con.Query(0,11,33)
	t.Logf("res:%+v", res)
}

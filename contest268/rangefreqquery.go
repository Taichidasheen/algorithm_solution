package contest268

/*
 https://leetcode.com/contest/weekly-contest-268/problems/range-frequency-queries/
 */
type RangeFreqQuery struct {
	arr []int
	exist map[int]int
}


func Constructor(arr []int) RangeFreqQuery {
	exist := make(map[int]int)
	for _, item := range arr {
		exist[item] = 1
	}
	return RangeFreqQuery{
		arr: arr,
		exist: exist,
	}
}


func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	var cnt int
	if this.exist[value] == 0 {
		return 0
	}
	for i:= left;i<=right;i++ {
		if this.arr[i] == value {
			cnt ++
		}
	}
	return cnt
}


/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
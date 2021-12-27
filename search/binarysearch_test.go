package search

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		left   int
		right  int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{-1,0,3,5,9,12},
				left: 0,
				right: 5,
				target: 2,
			},
			want: -1,
		},
		{
			args: args{
				nums: []int{5},
				left: 0,
				right: 0,
				target: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchV2(tt.args.nums, tt.args.left, tt.args.right, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
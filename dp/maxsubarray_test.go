package dp

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				nums: []int{-2,1,-3,4,-1,2,1,-5,4},
			},
			want: 6,
		},
		{
			name: "02",
			args: args{
				nums: []int{-2,1,-3,4},
			},
			want: 4,
		},
		{
			name: "03",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "04",
			args: args{
				nums: []int{-3,-2,-1},
			},
			want: -1,
		},
		{
			name: "05",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

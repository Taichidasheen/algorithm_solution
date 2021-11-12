package dp

import "testing"

func Test_lengthOfLIS(t *testing.T) {
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
				nums: []int{1,2,3,4},
			},
			want:4,
		},
		{
			name: "02",
			args: args{
				nums: []int{1,2,3,1},
			},
			want:3,
		},
		{
			name: "03",
			args: args{
				nums: []int{4,3,2,1},
			},
			want:1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

package contest268

import "testing"

func Test_binarySearchLowBound(t *testing.T) {
	type args struct {
		arr   []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{2,4,6,8,10},
				value: 9,
			},
			want: 3,
		},
		{
			args: args{
				arr: []int{2,4,6,8,10},
				value: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchLowBound(tt.args.arr, tt.args.value); got != tt.want {
				t.Errorf("binarySearchLowBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchUpperBound(t *testing.T) {
	type args struct {
		arr   []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{2,4,6,8,10},
				value: 9,
			},
			want: 3,
		},
		{
			args: args{
				arr: []int{2,4,6,8,10},
				value: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchUpperBound(tt.args.arr, tt.args.value); got != tt.want {
				t.Errorf("binarySearchUpperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
package sort

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "123_456",
			args: args{
				arr1: []int{1,2,3},
				arr2: []int{4,5,6},
			},
			want: []int{1,2,3,4,5,6},
		},
		{
			name: "135_246",
			args: args{
				arr1: []int{1,3,5},
				arr2: []int{2,4,6},
			},
			want: []int{1,2,3,4,5,6},
		},
		{
			name: "0_246",
			args: args{
				arr1: []int{},
				arr2: []int{2,4,6},
			},
			want: []int{2,4,6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "01",
			args: args{
				arr: []int{9,8,7,6,5,4,3,2,1},
			},
			want: []int{1,2,3,4,5,6,7,8,9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
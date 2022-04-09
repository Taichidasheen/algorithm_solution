package common

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				x: 25,
			},
			want: 5,
		},
		{
			args: args{
				x: 24,
			},
			want: 4,
		},
		{
			args: args{
				x: 100,
			},
			want: 10,
		},
		{
			args: args{
				x: 101,
			},
			want: 10,
		},
		{
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			args: args{
				x: 2,
			},
			want: 1,
		},
		{
			args: args{
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

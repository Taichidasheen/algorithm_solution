package dp

import (
	"testing"
)

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				envelopes: [][]int{
					[]int{5,4},
					[]int{6,4},
					[]int{6,7},
					[]int{2,3},
				},
			},
			want:3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortEnvelops(t *testing.T) {

	envelops := [][]int{
		[]int{5,4},
		[]int{6,4},
		[]int{6,7},
		[]int{2,3},
		[]int{1,4},
	}

	res := sortEnvelops(envelops)
	t.Logf("res:%+v", res)
}
package slidingwindow

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				s:"ADOBECODEBANC",
				t:"ABC",
			},
			want: "BANC",
		},
		{
			name: "02",
			args: args{
				s:"a",
				t:"a",
			},
			want: "a",
		},
		{
			name: "03",
			args: args{
				s:"a",
				t:"aa",
			},
			want: "",
		},
		{
			name: "04",
			args: args{
				s:"aaa",
				t:"aa",
			},
			want: "aa",
		},
		{
			name: "05",
			args: args{
				s:"cabwefgewcwaefgcf",
				t:"cae",
			},
			want: "cwae",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

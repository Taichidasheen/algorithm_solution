package palindrome

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "02",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "03",
			args: args{
				s: "abc",
			},
			want: "a",
		},
		{
			name: "04",
			args: args{
				s: "aabc",
			},
			want: "aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

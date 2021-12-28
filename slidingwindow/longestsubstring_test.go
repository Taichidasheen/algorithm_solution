package slidingwindow

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				"abcabcbb",
			},
			want: 3,
		},
		{
			name: "02",
			args: args{
				"bbbbb",
			},
			want: 1,
		},
		{
			name: "03",
			args: args{
				"pwwkew",
			},
			want: 3,
		},
		{
			name: "04",
			args: args{
				"",
			},
			want: 0,
		},
		{
			name: "05",
			args: args{
				"a",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
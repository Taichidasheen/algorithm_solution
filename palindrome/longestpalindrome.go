package palindrome

/*
lc-5: https://leetcode.com/problems/longest-palindromic-substring/
 */
func longestPalindrome(s string) string {
	var res string
	for i:=0;i<len(s);i++ {
		res1 := palindrome(s, i,i)
		res2 := palindrome(s, i,i+1)
		if len(res1) > len(res) {
			res = res1
		}
		if len(res2) > len(res) {
			res = res2
		}
	}
	return res
}

func palindrome(s string, l, r int) string {
	for l>=0 && r<= len(s) -1 && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1: r-1 + 1]
}


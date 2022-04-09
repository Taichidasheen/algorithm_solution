package common

import "testing"

func Test_decodeString(t *testing.T) {
	/*s := "3[a2[c]]"
	res := decodeString(s)
	t.Logf("res:%+v", res)*/

	s := "3[a]2[bc]"
	res := decodeString(s)
	t.Logf("res:%+v", res)

	s = "2[abc]3[cd]ef"
	res = decodeString(s)
	t.Logf("res:%+v", res)

}

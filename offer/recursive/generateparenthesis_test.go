package recursive

import (
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	res := generateParenthesis(1)
	t.Logf("res:%+v", res)
}
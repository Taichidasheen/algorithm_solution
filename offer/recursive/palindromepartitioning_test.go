package recursive

import (
	"testing"
)


func Test_partition(t *testing.T) {
	res := partition("google")
	t.Logf("res:%+v", res)
	res = partition("aab")
	t.Logf("res:%+v", res)
	res = partition("a")
	t.Logf("res:%+v", res)
	res = partition("bb")
	t.Logf("res:%+v", res)
}
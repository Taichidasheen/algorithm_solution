package lru

import "testing"

func TestHello(t *testing.T) {
	lru := NewLinkedHashMap(3)
	lru.Put(1, 111)
	t.Logf("lru:%+v", lru)
	lru.Put(2, 222)
	t.Logf("lru:%+v", lru)
	lru.Put(3, 333)
	t.Logf("lru:%+v", lru)
	lru.Put(4, 444)
	t.Logf("lru:%+v", lru)
}
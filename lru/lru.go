package lru

/*
 lc-146: https://leetcode.com/problems/lru-cache/
 */

type LRUCache struct {
	lhm *LinkedHashMap
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		lhm: NewLinkedHashMap(capacity),
	}
}


func (this *LRUCache) Get(key int) int {
	return this.lhm.Get(key)
}


func (this *LRUCache) Put(key int, value int)  {
	this.lhm.Put(key, value)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

package practice

import "container/list"

type LRU struct {
	capacity int
	list *list.List
	hash map[string]*list.Element
}

type Node struct {
	Key string
	Value interface{}
	Prev *Node
	Next *Node
}
type List struct {
	Head *Node
	Tail *Node
}

func NewList() *List {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return &List{
		Head: head,
		Tail: tail,
	}
}

func (l *List) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		hash: make(map[string]*list.Element),
	}
}

func (l *LRU) Get(key string) interface{} {
	if node, exist := l.hash[key];exist {
		l.list.Remove(node)
		l.list.PushBack(node)
		return node.Value
	}
	return nil
}

func (l *LRU) Set(key string, val interface{}) {
	if len(l.hash) >= l.capacity {
	}
}

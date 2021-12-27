package lru


//实现双向链表

type Node struct {
	Key int
	Value int
	Prev, Next *Node
}

type List struct {
	root *Node
	len int
}

//NewList 创建链表
func NewList() *List {
	root := &Node{}
	root.Next = root
	root.Prev = root
	l :=  &List{
		root:root,
		len: 0,
	}
	return l
}

//PushBack 向尾部添加元素
func (l *List) PushBack(n *Node) {
	lastTail := l.root.Prev
	lastTail.Next = n
	n.Prev = lastTail
	l.root.Prev = n
	l.len++
}

//PushFront 向头部添加元素
func (l *List) PushFront(n *Node) {
	lastHead := l.root.Next
	l.root.Next = n
	n.Prev = l.root
	n.Next = lastHead
	lastHead.Prev = n
	l.len++
}

//PopBack 返回链表最后一个节点
func (l *List) PopBack() *Node {
	tail := l.root.Prev
	l.Remove(tail)
	return tail
}

func (l *List) PopFront() *Node {
	head := l.root.Next
	l.Remove(head)
	return head
}

//Remove 删除元素
func (l *List) Remove(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	l.len--
}


package practice


type LRU struct {
	capacity int
	list *List
	hash map[string]*Node
}

type Node struct {
	Key string
	Value int
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

//放到头部
func (l *List) Push(node *Node) {
	tmp := l.Head.Next
	l.Head.Next = node
	node.Prev = l.Head
	node.Next = tmp
	tmp.Prev = node
}

//淘汰尾部
func (l *List) DelLast() *Node {
	last := l.Tail.Prev
	last.Prev.Next = l.Tail
	l.Tail.Prev = last.Prev
	return last
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		hash: make(map[string]*Node),
	}
}

func (l *LRU) Get(key string) interface{} {
	if node, exist := l.hash[key];exist {
		l.list.Remove(node)
		l.list.Push(node)
		return node.Value
	}
	return nil
}

func (l *LRU) Set(key string, val int) {
	node := &Node{
		Key: key,
		Value: val,
	}
	//原来已经存在
	if tmp, exist := l.hash[key];exist {
		tmp.Value = val
		l.list.Remove(tmp)
		l.list.Push(tmp)
	}
	//原来不存在,容量已经满了
	if len(l.hash) >= l.capacity {
		//淘汰节点后再插入
		last := l.list.DelLast()
		l.hash[last.Key] = nil
		l.list.Push(node)
		l.hash[key] = node
	} else {
		//原来不存在，容量还没有满
		l.list.Push(node)
		l.hash[key] = node
	}

}

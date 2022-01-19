package lru

//LinkedHashMap map + 链表
type LinkedHashMap struct {
	size int
	hash map[int]*Node
	list *List
}

//NewLinkedHashMap 初始化
func NewLinkedHashMap(size int) *LinkedHashMap {
	return &LinkedHashMap{
		size: size,
		hash: make(map[int]*Node),
		list: NewList(),
	}
}

//Get 获取Value
func (m *LinkedHashMap) Get(key int) int {
	if node, ok := m.hash[key]; ok {
		m.list.Remove(node)
		m.list.PushBack(node)
		return node.Value
	}
	return -1
}

func (m *LinkedHashMap) Put(key, val int) {
	if node, ok := m.hash[key]; ok {
		//update
		m.list.Remove(node)
		node.Value = val
		m.list.PushBack(node)
	} else {
		//add
		nod := &Node{
			Key: key,
			Value: val,
		}
		if len(m.hash)>=m.size {
			front := m.list.PopFront()//删掉
			delete(m.hash, front.Key)
			m.list.PushBack(nod)
			//不要忘记把key加到map中
			m.hash[key] = nod
		} else {
			m.list.PushBack(nod)
			//不要忘记把key加到map中
			m.hash[key] = nod
		}
	}
}


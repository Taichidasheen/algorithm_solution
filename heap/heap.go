package heap

/*
 堆的相关知识可以参考：
https://my.oschina.net/xinxingegeya/blog/703801
https://my.oschina.net/xinxingegeya/blog/705409

 */

type Heap struct {
	nodes []int
	size int
}

func NewHeap() *Heap {
	return &Heap{
		size: 0,
		nodes: make([]int, 1),//nodes[0]空出
	}
}

func (h *Heap) Init() {
	for i:=h.size/2;i>=1;i-- {
		h.down(i)
	}
}

//Pop 取出堆顶元素
// Pop的核心：
// 1. 取出第一个元素后，把最后一个元素放到第一个位置做一次向下调整
// 2. 去掉一个元素后，记得更新slice的长度
func(h *Heap) Pop() int {
	if h.size > 0 {
		top := h.nodes[1]
		h.nodes[1] = h.nodes[h.size]
		h.nodes = h.nodes[0:h.size]//去掉最后一个元素
		h.size-- //size减小
		h.down(1)
		return top
	}
	//假设堆中的元素都大于0
	return -1
}

//Push 往堆中插入元素
func(h *Heap) Push(val int) {
	h.nodes = append(h.nodes, val)
	h.size++
	h.up(h.size)
}

//up 向上调整
// 向上调整的核心：（下标从1开始的情况）
// 1. 父节点下标为i/2
func(h *Heap) up(i int) {
	ch := i
	for ch > 1 {
		pa := ch/2
		if h.nodes[ch] > h.nodes[pa] {
			break
		} else {
			h.swap(ch, pa)
			ch = pa
		}
	}
}

//down 向下调整
// 向下调整的核心：（下标从1开始的情况）
// 1. 左子节点下标为2i，右子节点下标为2i+1
// 2. 父节点需要和子节点中较小的那一个进行互换
func(h *Heap) down(i int) {
	pa := i
	for pa < h.size {
		left := pa*2
		if left > h.size {
			break
		}
		var smaller int
		right := pa*2 + 1
		if right > h.size {
			smaller = left
		} else {
			if h.nodes[left] < h.nodes[right] {
				smaller = left
			} else {
				smaller = right
			}
		}
		if h.nodes[pa] > h.nodes[smaller] {
			h.swap(pa, smaller)
			pa = smaller
		} else {
			break
		}
	}
}

func(h *Heap) swap(i, j int) {
	temp := h.nodes[i]
	h.nodes[i] = h.nodes[j];
	h.nodes[j] = temp
}

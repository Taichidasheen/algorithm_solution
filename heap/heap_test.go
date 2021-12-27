package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap()
	heap.Init()
	t.Logf("heap:%+v", heap)
	heap.Push(5)
	heap.Push(4)
	heap.Push(3)
	heap.Push(2)
	heap.Push(6)
	t.Logf("heap:%+v", heap)
	heap.Pop()
	t.Logf("heap:%+v", heap)
	heap.Pop()
	t.Logf("heap:%+v", heap)
}
package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	minHeap := NewMinHeap()
	minHeap.Insert(2)
	minHeap.Insert(5)
	minHeap.Insert(1)
	minHeap.Insert(6)
	minHeap.Insert(7)

	fmt.Println(minHeap.String())


	sysHeap := make(SysHeap, 0)
	heap.Init(&sysHeap)

	heap.Push(&sysHeap, 2)
	heap.Push(&sysHeap, 5)
	heap.Push(&sysHeap, 1)
	heap.Push(&sysHeap, 6)
	heap.Push(&sysHeap, 7)

	fmt.Println(heap.Pop(&sysHeap))
	fmt.Println(heap.Pop(&sysHeap))
	fmt.Println(heap.Pop(&sysHeap))
	fmt.Println(heap.Pop(&sysHeap))
	fmt.Println(heap.Pop(&sysHeap))
}
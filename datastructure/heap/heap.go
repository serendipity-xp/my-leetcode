package heap

import (
	"fmt"
	"math"
)

type Heap struct {
	Element []int
}

func NewMinHeap() *Heap {
	return &Heap{Element: []int{math.MinInt64}}
}

func (m *Heap) Insert(val int) {
	m.Element = append(m.Element, val)
	i := len(m.Element) - 1
	for ; m.Element[i/2] > val; i /= 2 {
		m.Element[i] = m.Element[i/2]
	}
	m.Element[i] = val
}

func (m *Heap) DeleteMin() (int, error) {
	if len(m.Element) <= 1 {
		return 0, fmt.Errorf("heap is empty")
	}
	maxElement := m.Element[0]
	lastElement := m.Element[len(m.Element)-1]
	var i, child int
	for i = 1; i*2 < len(m.Element); i = child {
		child = i * 2
		if child < len(m.Element)-1 && m.Element[child+1] < m.Element[child] {
			child++
		}
		if lastElement > m.Element[child] {
			m.Element[i] = m.Element[child]
		} else {
			break
		}
	}
	m.Element[i] = lastElement
	m.Element = m.Element[:len(m.Element)-1]
	return maxElement, nil

}

func (m *Heap) Length() int {
	return len(m.Element)
}

func (m *Heap) Min() (int, error) {
	if len(m.Element) > 1{
		return m.Element[1], nil
	}
	return 0, fmt.Errorf("heap is empty")
}

func (m *Heap) String() string {
	return fmt.Sprintf("%v", m.Element)
}



type SysHeap []int

func (h SysHeap) Len() int {
	return len(h)
}

func (h SysHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *SysHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *SysHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *SysHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return result
}

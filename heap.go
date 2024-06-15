package datastruct

import (
	"sync"
)

type Heap[T any] struct {
	desc  bool
	items []heapNode[T]
	mx    sync.RWMutex
}

type heapNode[T any] struct {
	priority int64
	value    T
}

func NewHeap[T any](desc bool, cap int) *Heap[T] {
	return &Heap[T]{
		desc:  desc,
		items: make([]heapNode[T], 0, cap),
	}
}

func (h *Heap[T]) Push(priority int64, val T) {
	h.mx.Lock()
	defer h.mx.Unlock()
	h.items = append(h.items, heapNode[T]{
		priority: priority,
		value:    val,
	})
	h.heapifyUp()
}

func (h *Heap[T]) Pop() (val T, ok bool) {
	h.mx.Lock()
	defer h.mx.Unlock()
	if len(h.items) == 0 {
		return val, false
	}
	if len(h.items) == 1 {
		val, h.items = h.items[0].value, h.items[:0]
		return val, true
	}
	val = h.items[0].value
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.heapifyDown()
	return val, true
}

func (h *Heap[T]) heapifyDown() {
	curr := 0
	for curr < len(h.items) {
		left := h.left(curr)
		right := h.right(curr)
		child := curr
		if h.desc {
			child = h.larger(h.larger(curr, left), right)
		} else {
			child = h.smaller(h.smaller(curr, left), right)
		}
		if child > curr {
			h.swap(curr, child)
			curr = child
			continue
		}
		break
	}
}

func (h *Heap[T]) heapifyUp() {
	curr := len(h.items) - 1
	for curr > 0 {
		parent := h.parent(curr)
		if h.desc {
			if h.items[curr].priority > h.items[parent].priority {
				// fmt.Println("swap:", h.items[parent].priority, h.items[curr].priority)
				h.swap(parent, curr)
				curr = parent
				continue
			}
		} else {
			if h.items[curr].priority < h.items[parent].priority {
				// fmt.Println("swap:", h.items[parent].priority, h.items[curr].priority)
				h.swap(parent, curr)
				curr = parent
				continue
			}
		}
		break
	}
}

func (h *Heap[T]) parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap[T]) left(index int) int {
	res := 2*index + 1
	if res > len(h.items)-1 {
		return -1
	}
	return res
}

func (h *Heap[T]) right(index int) int {
	res := 2*index + 2
	if res > len(h.items)-1 {
		return -1
	}
	return res
}

func (h *Heap[T]) swap(a, b int) {
	h.items[a], h.items[b] = h.items[b], h.items[a]
}

func (h *Heap[T]) larger(a, b int) int {
	if a < 0 && b < 0 {
		return -1
	}
	if a < 0 && b > -1 {
		return b
	}
	if a > -1 && b < 0 {
		return a
	}
	if h.items[a].priority > h.items[b].priority {
		return a
	}
	return b
}

func (h *Heap[T]) smaller(a, b int) int {
	if a < 0 && b < 0 {
		return -1
	}
	if a < 0 && b > -1 {
		return b
	}
	if a > -1 && b < 0 {
		return a
	}
	if h.items[a].priority < h.items[b].priority {
		return a
	}
	return b
}

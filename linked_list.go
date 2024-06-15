package datastruct

import (
	"sync"
	"sync/atomic"
)

type LinkedList[T any] struct {
	sync.Mutex
	next *LNode[T]
	size int64
}

type LNode[T any] struct {
	next  *LNode[T]
	value T
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (list *LinkedList[T]) Len() int64 {
	return atomic.LoadInt64(&list.size)
}

func (list *LinkedList[T]) Push(val T) {
	newNode := &LNode[T]{
		value: val,
	}
	list.Lock()
	newNode.next, list.next = list.next, newNode
	list.Unlock()
	atomic.AddInt64(&list.size, 1)
}

func (list *LinkedList[T]) Pop() (val T, ok bool) {
	list.Lock()
	if list.next == nil {
		list.Unlock()
		return val, false
	}
	val, list.next = list.next.value, list.next.next
	list.Unlock()
	atomic.AddInt64(&list.size, -1)
	return val, true
}

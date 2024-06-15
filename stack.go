package datastruct

type Stack[T any] struct {
	cap  int64
	list *LinkedList[T]
}

func NewStack[T any](capacity int64) *Stack[T] {
	return &Stack[T]{
		cap:  capacity,
		list: NewLinkedList[T](),
	}
}

func (this *Stack[T]) Len() int64 {
	return this.list.Len()
}

func (this *Stack[T]) Push(val T) bool {
	if this.Len() >= this.cap {
		return false
	}
	this.list.Push(val)
	return true
}

func (this *Stack[T]) Pop() (val T, ok bool) {
	return this.list.Pop()
}

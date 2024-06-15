package datastruct

type Queue[T any] struct {
	items chan T
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		items: make(chan T, size),
	}
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}

func (q *Queue[T]) Push(val T) bool {
	select {
	case q.items <- val:
		return true
	default:
		return false
	}
}

func (q *Queue[T]) Pop() (val T, ok bool) {
	select {
	case val := <-q.items:
		return val, true
	default:
		return val, false
	}
}

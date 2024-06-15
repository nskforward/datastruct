package datastruct

import (
	"math/rand"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int](8)
	_, ok := q.Pop()
	if ok {
		t.Fatalf("must be empty")
	}
	q.Push(1)
	v, ok := q.Pop()
	if !ok {
		t.Fatalf("must have value")
	}
	if v != 1 {
		t.Fatalf("value must be 1, actual %d", v)
	}
	q.Push(0)
	q.Push(1)
	q.Push(2)
	for i := 0; i < 3; i++ {
		v, ok := q.Pop()
		if !ok {
			t.Fatalf("index %d must have value", i)
		}
		if v != i {
			t.Fatalf("index %d value must be %d, actual %d", i, i, v)
		}
	}
	_, ok = q.Pop()
	if ok {
		t.Fatalf("must be empty")
	}
}

// 278.0 ns/op
func BenchmarkQueue(b *testing.B) {
	s := NewQueue[int](32)
	for i := 0; i < b.N; i++ {
		l := s.Len()
		if rand.Intn(2) == 0 && l < 32 {
			s.Push(1)
			continue
		}
		if l > 0 {
			val, ok := s.Pop()
			if ok {
				_ = val
			}
		}
	}
}

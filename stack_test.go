package datastruct

import (
	"math/rand"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int](64)
	_, ok := s.Pop()
	if ok {
		t.Fatalf("must be empty")
	}
	s.Push(1)
	v, ok := s.Pop()
	if !ok {
		t.Fatalf("must have value")
	}
	if v != 1 {
		t.Fatalf("value must be 1, actual %d", v)
	}
	s.Push(2)
	s.Push(1)
	s.Push(0)
	for i := 0; i < 3; i++ {
		v, ok := s.Pop()
		if !ok {
			t.Fatalf("index %d must have value", i)
		}
		if v != i {
			t.Fatalf("index %d value must be %d, actual %d", i, i, v)
		}
	}
	_, ok = s.Pop()
	if ok {
		t.Fatalf("must be empty")
	}
}

// 1058/2 ns/op
func BenchmarkStack(b *testing.B) {
	s := NewStack[int](64)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			l := s.Len()
			if rand.Intn(2) == 0 && l < 64 {
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
	})
}

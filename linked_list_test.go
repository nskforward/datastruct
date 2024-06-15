package datastruct

import (
	"fmt"
	"testing"
)

func TestLinkedListSingleThread(t *testing.T) {
	ll := NewLinkedList[int]()
	for i := 0; i < 10; i++ {
		ll.Push(i)
	}
	fmt.Println("size before:", ll.Len())
	for i := 0; i < 10; i++ {
		val, ok := ll.Pop()
		if !ok {
			t.Fatalf("value must exist")
		}
		expect := 9 - i
		if val != expect {
			t.Fatalf("expect value %d, actual %d", expect, val)
		}
	}
	fmt.Println("size after:", ll.Len())
}

// 1137/2 ns/op
func BenchmarkLinkedListPushAndPopMultiThread(b *testing.B) {
	ll := NewLinkedList[bool]()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			ll.Push(true)
			ll.Pop()
		}
	})
}

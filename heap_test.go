package datastruct

import (
	"fmt"
	"testing"
)

func TestHeapMax(t *testing.T) {
	testCases := [2][]int{
		{1, 3, 2, 7, 4, 5, 6, 8}, {8, 7, 6, 5, 4, 3, 2, 1},
	}
	h := NewHeap[int](true, 8)
	for _, tc := range testCases[0] {
		h.Push(int64(tc), tc)
	}
	for _, tc := range testCases[1] {
		val, ok := h.Pop()
		if !ok {
			t.Fatalf("value must exist for test case %d", tc)
		}
		if val != tc {
			t.Fatalf("expect %d, actual %d", tc, val)
		}
		fmt.Print(val)
	}
	fmt.Println()
}

func TestHeapMin(t *testing.T) {
	testCases := [2][]int{
		{1, 3, 2, 7, 4, 5, 6, 8}, {1, 2, 3, 4, 5, 6, 7, 8},
	}
	h := NewHeap[int](false, 8)
	for _, tc := range testCases[0] {
		h.Push(int64(tc), tc)
	}
	fmt.Println(h.items)
	for _, tc := range testCases[1] {
		val, ok := h.Pop()
		if !ok {
			t.Fatalf("value must exist for test case %d", tc)
		}
		if val != tc {
			t.Fatalf("expect %d, actual %d", tc, val)
		}
		fmt.Print(val)
	}
	fmt.Println()
}

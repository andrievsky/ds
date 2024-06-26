package heap

import (
	"testing"
)

func TestPairingHeapPeek(t *testing.T) {
	heap := NewMaxHeap[int]()
	heap.Push(1)
	actual := heap.Peek()
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingHeapPeek2(t *testing.T) {
	heap := NewMaxHeap[int]()
	heap.Push(1)
	heap.Push(2)
	actual := heap.Peek()
	expected := 2
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingHeapPop(t *testing.T) {
	heap := NewMaxHeap[int]()
	heap.Push(1)
	actual := heap.Pop()
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingHeapPop2(t *testing.T) {
	heap := NewMaxHeap[int]()
	heap.Push(1)
	heap.Push(2)
	actual := heap.Pop()
	expected := 2
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	actual = heap.Pop()
	expected = 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingHeapIsEmpty(t *testing.T) {
	heap := NewMaxHeap[int]()
	actual := heap.IsEmpty()
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingHeapIsEmpty2(t *testing.T) {
	heap := NewMaxHeap[int]()
	heap.Push(1)
	actual := heap.IsEmpty()
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	heap.Pop()
	actual = heap.IsEmpty()
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

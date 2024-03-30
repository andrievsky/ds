package ph

import (
	"testing"
)

func TestPairingTreePeek(t *testing.T) {
	heap := Insert(nil, 1)
	actual := Peek(heap)
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingTreePeek2(t *testing.T) {
	heap := Insert(nil, 1)
	heap = Insert(heap, 2)
	actual := Peek(heap)
	expected := 2
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingTreePop(t *testing.T) {
	heap := Insert(nil, 1)
	actual, _ := Pop(heap)
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingTreePop2(t *testing.T) {
	heap := Insert(nil, 1)
	heap = Insert(heap, 2)
	actual := 0
	actual, heap = Pop(heap)
	expected := 2
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	actual, _ = Pop(heap)
	expected = 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingTreeIsEmpty(t *testing.T) {
	heap := Insert(nil, 1)
	actual := IsEmpty(heap)
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPairingTreeIsEmpty2(t *testing.T) {
	heap := Insert(nil, 1)
	heap = Insert(heap, 2)
	actual := IsEmpty(heap)
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	_, heap = Pop(heap)
	_, heap = Pop(heap)
	actual = IsEmpty(heap)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

package stack

import "testing"

func TestBuildStack(t *testing.T) {
	s := BuildStack[int](1, 2, 3)
	if s == nil {
		t.Errorf("BuildStack() returned nil")
	}
}

func TestStack_Push(t *testing.T) {
	s := BuildStack[int]()
	s.Push(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("Stack.Push() failed")
	}
}

func TestStack_Pop(t *testing.T) {
	s := BuildStack[int](1, 2, 3)
	if s.Pop() != 3 {
		t.Errorf("Stack.Pop() failed")
	}
}

func TestStack_Peek(t *testing.T) {
	s := BuildStack[int](1, 2, 3)
	if s.Peek() != 3 {
		t.Errorf("Stack.Peek() failed")
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := BuildStack[int]()
	if !s.IsEmpty() {
		t.Errorf("Stack.IsEmpty() failed")
	}
}

func TestStack_Size(t *testing.T) {
	s := BuildStack[int](1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("Stack.Size() failed")
	}
}

func TestStack_Clear(t *testing.T) {
	s := BuildStack[int](1, 2, 3)
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("Stack.Clear() failed")
	}
}

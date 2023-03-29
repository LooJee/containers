package stack

import "testing"

func TestBuildThreadsafeStack(t *testing.T) {
	s := BuildThreadsafeStack[int](1, 2, 3)
	if s == nil {
		t.Errorf("BuildThreadsafeStack() returned nil")
	}
}

func TestThreadsafeStack_Push(t *testing.T) {
	s := BuildThreadsafeStack[int]()
	s.Push(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("ThreadsafeStack.Push() failed")
	}
}

func TestThreadsafeStack_Pop(t *testing.T) {
	s := BuildThreadsafeStack[int](1, 2, 3)
	if s.Pop() != 3 {
		t.Errorf("ThreadsafeStack.Pop() failed")
	}
}

func TestThreadsafeStack_Peek(t *testing.T) {
	s := BuildThreadsafeStack[int](1, 2, 3)
	if s.Peek() != 3 {
		t.Errorf("ThreadsafeStack.Peek() failed")
	}
}

func TestThreadsafeStack_IsEmpty(t *testing.T) {
	s := BuildThreadsafeStack[int]()
	if !s.IsEmpty() {
		t.Errorf("ThreadsafeStack.IsEmpty() failed")
	}
}

func TestThreadsafeStack_Size(t *testing.T) {
	s := BuildThreadsafeStack[int](1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("ThreadsafeStack.Size() failed")
	}
}

func TestThreadsafeStack_Clear(t *testing.T) {
	s := BuildThreadsafeStack[int](1, 2, 3)
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("ThreadsafeStack.Clear() failed")
	}
}

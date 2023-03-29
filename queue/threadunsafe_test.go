package queue

import "testing"

func TestBuildQueue(t *testing.T) {
	q := BuildQueue[int](1, 2, 3)
	if q == nil {
		t.Errorf("BuildQueue() returned nil")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	q := BuildQueue[int]()
	q.Enqueue(1, 2, 3)
	if q.Size() != 3 {
		t.Errorf("Queue.Enqueue() failed")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := BuildQueue[int](1, 2, 3)
	if q.Dequeue() != 1 {
		t.Errorf("Queue.Dequeue() failed")
	}
}

func TestQueue_Peek(t *testing.T) {
	q := BuildQueue[int](1, 2, 3)
	if q.Peek() != 1 {
		t.Errorf("Queue.Peek() failed")
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := BuildQueue[int]()
	if !q.IsEmpty() {
		t.Errorf("Queue.IsEmpty() failed")
	}
}

func TestQueue_Size(t *testing.T) {
	q := BuildQueue[int](1, 2, 3)
	if q.Size() != 3 {
		t.Errorf("Queue.Size() failed")
	}
}

func TestQueue_Clear(t *testing.T) {
	q := BuildQueue[int](1, 2, 3)
	q.Clear()
	if q.Size() != 0 {
		t.Errorf("Queue.Clear() failed")
	}
}

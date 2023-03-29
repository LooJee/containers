package queue

import "testing"

func TestBuildThreadsafeQueue(t *testing.T) {
	q := BuildThreadsafeQueue[int](1, 2, 3)
	if q == nil {
		t.Errorf("BuildThreadsafeQueue() returned nil")
	}
}

func TestThreadsafeQueue_Enqueue(t *testing.T) {
	q := BuildThreadsafeQueue[int]()
	q.Enqueue(1, 2, 3)
	if q.Size() != 3 {
		t.Errorf("ThreadsafeQueue.Enqueue() failed")
	}
}

func TestThreadsafeQueue_Dequeue(t *testing.T) {
	q := BuildThreadsafeQueue[int](1, 2, 3)
	if q.Dequeue() != 1 {
		t.Errorf("ThreadsafeQueue.Dequeue() failed")
	}
}

func TestThreadsafeQueue_Peek(t *testing.T) {
	q := BuildThreadsafeQueue[int](1, 2, 3)
	if q.Peek() != 1 {
		t.Errorf("ThreadsafeQueue.Peek() failed")
	}
}

func TestThreadsafeQueue_IsEmpty(t *testing.T) {
	q := BuildThreadsafeQueue[int]()
	if !q.IsEmpty() {
		t.Errorf("ThreadsafeQueue.IsEmpty() failed")
	}
}

func TestThreadsafeQueue_Size(t *testing.T) {
	q := BuildThreadsafeQueue[int](1, 2, 3)
	if q.Size() != 3 {
		t.Errorf("ThreadsafeQueue.Size() failed")
	}
}

func TestThreadsafeQueue_Clear(t *testing.T) {
	q := BuildThreadsafeQueue[int](1, 2, 3)
	q.Clear()
	if q.Size() != 0 {
		t.Errorf("ThreadsafeQueue.Clear() failed")
	}
}

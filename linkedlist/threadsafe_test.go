package linkedlist

import "testing"

func TestThreadsafeLinkedList_Append(t *testing.T) {
	l := BuildThreadsafeLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}

	if l.Size() != 3 {
		t.Errorf("ThreadsafeLinkedList.Append() failed")
	}
}

func TestThreadsafeLinkedList_Prepend(t *testing.T) {
	l := BuildThreadsafeLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Prepend(data)
	}
	if l.Size() != 3 {
		t.Errorf("ThreadsafeLinkedList.Prepend() failed")
	}
}

func TestThreadsafeLinkedList_Remove(t *testing.T) {
	l := BuildThreadsafeLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.Remove(1)
	if l.Size() != 2 {
		t.Errorf("ThreadsafeLinkedList.Remove() failed")
	}
}

func TestThreadsafeLinkedList_RemoveAt(t *testing.T) {
	l := BuildThreadsafeLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.RemoveAt(1)
	if l.Size() != 2 {
		t.Errorf("ThreadsafeLinkedList.RemoveAt() failed")
	}
}

func TestThreadsafeLinkedList_IsEmpty(t *testing.T) {
	l := BuildThreadsafeLinkedList[int]()
	if !l.IsEmpty() {
		t.Errorf("ThreadsafeLinkedList.IsEmpty() failed")
	}
}

func TestThreadsafeLinkedList_Size(t *testing.T) {
	l := BuildThreadsafeLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.Size() != 3 {
		t.Errorf("ThreadsafeLinkedList.Size() failed")
	}
}

func TestThreadsafeLinkedList_Contains(t *testing.T) {
	l := BuildThreadsafeLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if !l.Contains(1) {
		t.Errorf("ThreadsafeLinkedList.Contains() failed")
	}
}

func TestThreadsafeLinkedList_Peek(t *testing.T) {
	l := BuildThreadsafeLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.Peek() != 1 {
		t.Errorf("ThreadsafeLinkedList.Peek() failed")
	}
}

func TestThreadsafeLinkedList_PeekLast(t *testing.T) {
	l := BuildThreadsafeLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.PeekLast() != 3 {
		t.Errorf("ThreadsafeLinkedList.PeekLast() failed")
	}
}

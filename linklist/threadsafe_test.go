package linklist

import "testing"

func TestThreadsafeLinkList_Append(t *testing.T) {
	l := BuildThreadsafeLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}

	if l.Size() != 3 {
		t.Errorf("ThreadsafeLinkList.Append() failed")
	}
}

func TestThreadsafeLinkList_Prepend(t *testing.T) {
	l := BuildThreadsafeLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Prepend(data)
	}
	if l.Size() != 3 {
		t.Errorf("ThreadsafeLinkList.Prepend() failed")
	}
}

func TestThreadsafeLinkList_Remove(t *testing.T) {
	l := BuildThreadsafeLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.Remove(1)
	if l.Size() != 2 {
		t.Errorf("ThreadsafeLinkList.Remove() failed")
	}
}

func TestThreadsafeLinkList_RemoveAt(t *testing.T) {
	l := BuildThreadsafeLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.RemoveAt(1)
	if l.Size() != 2 {
		t.Errorf("ThreadsafeLinkList.RemoveAt() failed")
	}
}

func TestThreadsafeLinkList_IsEmpty(t *testing.T) {
	l := BuildThreadsafeLinkList[int]()
	if !l.IsEmpty() {
		t.Errorf("ThreadsafeLinkList.IsEmpty() failed")
	}
}

func TestThreadsafeLinkList_Size(t *testing.T) {
	l := BuildThreadsafeLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.Size() != 3 {
		t.Errorf("ThreadsafeLinkList.Size() failed")
	}
}

func TestThreadsafeLinkList_Contains(t *testing.T) {
	l := BuildThreadsafeLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if !l.Contains(1) {
		t.Errorf("ThreadsafeLinkList.Contains() failed")
	}
}

func TestThreadsafeLinkList_Peek(t *testing.T) {
	l := BuildThreadsafeLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.Peek() != 1 {
		t.Errorf("ThreadsafeLinkList.Peek() failed")
	}
}

func TestThreadsafeLinkList_PeekLast(t *testing.T) {
	l := BuildThreadsafeLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.PeekLast() != 3 {
		t.Errorf("ThreadsafeLinkList.PeekLast() failed")
	}
}

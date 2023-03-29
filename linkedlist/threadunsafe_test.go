package linkedlist

import "testing"

func TestLinkedList_Append(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}

	if l.Size() != 3 {
		t.Errorf("LinkedList.Append() failed")
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Prepend(data)
	}
	if l.Size() != 3 {
		t.Errorf("LinkedList.Prepend() failed")
	}
}

func TestLinkedList_Remove(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.Remove(1)
	if l.Size() != 2 {
		t.Errorf("LinkedList.Remove() failed")
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.RemoveAt(1)
	if l.Size() != 2 {
		t.Errorf("LinkedList.RemoveAt() failed")
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	l := BuildLinkedList[int]()
	if !l.IsEmpty() {
		t.Errorf("LinkedList.IsEmpty() failed")
	}
}

func TestLinkedList_Size(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.Size() != 3 {
		t.Errorf("LinkedList.Size() failed")
	}
}

func TestLinkedList_Clear(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.Clear()
	if l.Size() != 0 {
		t.Errorf("LinkedList.Clear() failed")
	}
}

func TestLinkedList_Contains(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if !l.Contains(1) {
		t.Errorf("LinkedList.Contains() failed")
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.IndexOf(1) != 0 {
		t.Errorf("LinkedList.IndexOf() failed")
	}
}

func TestLinkedList_Peek(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.Peek() != 1 {
		t.Errorf("LinkedList.Peek() failed")
	}
}

func TestLinkedList_PeekLast(t *testing.T) {
	l := BuildLinkedList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.PeekLast() != 3 {
		t.Errorf("LinkedList.PeekLast() failed")
	}
}

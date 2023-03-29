package linklist

import "testing"

func TestLinkList_Append(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}

	if l.Size() != 3 {
		t.Errorf("LinkList.Append() failed")
	}
}

func TestLinkList_Prepend(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Prepend(data)
	}
	if l.Size() != 3 {
		t.Errorf("LinkList.Prepend() failed")
	}
}

func TestLinkList_Remove(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.Remove(1)
	if l.Size() != 2 {
		t.Errorf("LinkList.Remove() failed")
	}
}

func TestLinkList_RemoveAt(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.RemoveAt(1)
	if l.Size() != 2 {
		t.Errorf("LinkList.RemoveAt() failed")
	}
}

func TestLinkList_IsEmpty(t *testing.T) {
	l := BuildLinkList[int]()
	if !l.IsEmpty() {
		t.Errorf("LinkList.IsEmpty() failed")
	}
}

func TestLinkList_Size(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.Size() != 3 {
		t.Errorf("LinkList.Size() failed")
	}
}

func TestLinkList_Clear(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	l.Clear()
	if l.Size() != 0 {
		t.Errorf("LinkList.Clear() failed")
	}
}

func TestLinkList_Contains(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if !l.Contains(1) {
		t.Errorf("LinkList.Contains() failed")
	}
}

func TestLinkList_IndexOf(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.IndexOf(1) != 0 {
		t.Errorf("LinkList.IndexOf() failed")
	}
}

func TestLinkList_Peek(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.Peek() != 1 {
		t.Errorf("LinkList.Peek() failed")
	}
}

func TestLinkList_PeekLast(t *testing.T) {
	l := BuildLinkList[int]()
	for _, data := range []int{1, 2, 3} {
		l.Append(data)
	}
	if l.PeekLast() != 3 {
		t.Errorf("LinkList.PeekLast() failed")
	}
}

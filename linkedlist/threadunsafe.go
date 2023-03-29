package linkedlist

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type Node[T comparable] struct {
	data T
	next *Node[T]
}

// BuildLinkedList returns a new LinkedList[T] with the specified data.
func BuildLinkedList[T comparable]() *LinkedList[T] {
	list := &LinkedList[T]{head: nil, tail: nil, size: 0}
	return list
}

// Append appends the specified element to the end of this list.
func (l *LinkedList[T]) Append(data T) {
	node := &Node[T]{data: data, next: nil}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.size++
}

// Prepend adds data to the front of the list.
func (l *LinkedList[T]) Prepend(data T) {
	node := &Node[T]{data: data, next: nil}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.size++
}

// Remove remove the specified element from this list, if it is present.
func (l *LinkedList[T]) Remove(data T) {
	if l.head == nil {
		return
	}

	if l.head.data == data {
		l.head = l.head.next
		l.size--
		return
	}

	prev := l.head
	for prev.next != nil {
		if prev.next.data == data {
			prev.next = prev.next.next
			l.size--
			return
		}
		prev = prev.next
	}
}

// RemoveAt removes the element at the specified position in this list.
func (l *LinkedList[T]) RemoveAt(index int) {
	if index < 0 || index >= l.size {
		return
	}

	if index == 0 {
		l.head = l.head.next
		l.size--
		return
	}

	prev := l.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	l.size--
}

// InsertAt inserts the specified element at the specified position in this list.
func (l *LinkedList[T]) InsertAt(index int, data T) {
	if index < 0 || index >= l.size {
		return
	}

	if index == 0 {
		l.Prepend(data)
		return
	}

	node := &Node[T]{data: data, next: nil}
	prev := l.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}
	node.next = prev.next
	prev.next = node
	l.size++
}

// IndexOf returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (l *LinkedList[T]) IndexOf(data T) int {
	if l.head == nil {
		return -1
	}

	index := 0
	for node := l.head; node != nil; node = node.next {
		if node.data == data {
			return index
		}
		index++
	}

	return -1
}

// Contains returns true if the list contains the specified element.
func (l *LinkedList[T]) Contains(data T) bool {
	return l.IndexOf(data) != -1
}

// IsEmpty returns true if the list is empty.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

// Size returns the number of elements in the list.
func (l *LinkedList[T]) Size() int {
	return l.size
}

// Clear removes all elements from the list.
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList[T]) Peek() T {
	return l.head.data
}

func (l *LinkedList[T]) PeekLast() T {
	return l.tail.data
}

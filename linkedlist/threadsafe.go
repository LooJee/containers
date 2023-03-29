package linkedlist

import "sync"

type ThreadsafeLinkedList[T comparable] struct {
	lock sync.RWMutex
	data *LinkedList[T]
}

func BuildThreadsafeLinkedList[T comparable]() *ThreadsafeLinkedList[T] {
	return &ThreadsafeLinkedList[T]{
		data: BuildLinkedList[T](),
		lock: sync.RWMutex{},
	}
}

func (l *ThreadsafeLinkedList[T]) Append(data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.Append(data)
}

func (l *ThreadsafeLinkedList[T]) Prepend(data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.Prepend(data)
}

func (l *ThreadsafeLinkedList[T]) InsertAt(index int, data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.InsertAt(index, data)
}

func (l *ThreadsafeLinkedList[T]) RemoveAt(index int) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.RemoveAt(index)
}

func (l *ThreadsafeLinkedList[T]) Remove(data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.Remove(data)
}

func (l *ThreadsafeLinkedList[T]) IsEmpty() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.IsEmpty()
}

func (l *ThreadsafeLinkedList[T]) Size() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.Size()
}

func (l *ThreadsafeLinkedList[T]) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.Clear()
}

func (l *ThreadsafeLinkedList[T]) Contains(data T) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.Contains(data)
}

func (l *ThreadsafeLinkedList[T]) Peek() T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.Peek()
}

func (l *ThreadsafeLinkedList[T]) PeekLast() T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.PeekLast()
}

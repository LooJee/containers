package linklist

import "sync"

type ThreadsafeLinkList[T comparable] struct {
	lock sync.RWMutex
	data *LinkList[T]
}

func BuildThreadsafeLinkList[T comparable]() *ThreadsafeLinkList[T] {
	return &ThreadsafeLinkList[T]{
		data: BuildLinkList[T](),
		lock: sync.RWMutex{},
	}
}

func (l *ThreadsafeLinkList[T]) Append(data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.Append(data)
}

func (l *ThreadsafeLinkList[T]) Prepend(data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.Prepend(data)
}

func (l *ThreadsafeLinkList[T]) InsertAt(index int, data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.InsertAt(index, data)
}

func (l *ThreadsafeLinkList[T]) RemoveAt(index int) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.RemoveAt(index)
}

func (l *ThreadsafeLinkList[T]) Remove(data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.Remove(data)
}

func (l *ThreadsafeLinkList[T]) IsEmpty() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.IsEmpty()
}

func (l *ThreadsafeLinkList[T]) Size() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.Size()
}

func (l *ThreadsafeLinkList[T]) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data.Clear()
}

func (l *ThreadsafeLinkList[T]) Contains(data T) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.Contains(data)
}

func (l *ThreadsafeLinkList[T]) Peek() T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.Peek()
}

func (l *ThreadsafeLinkList[T]) PeekLast() T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.data.PeekLast()
}

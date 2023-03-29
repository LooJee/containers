package queue

import "sync"

type ThreadsafeQueue[T comparable] struct {
	lock sync.RWMutex
	data *Queue[T]
}

func BuildThreadsafeQueue[T comparable](data ...T) *ThreadsafeQueue[T] {
	return &ThreadsafeQueue[T]{
		data: BuildQueue[T](data...),
		lock: sync.RWMutex{},
	}
}

func (q *ThreadsafeQueue[T]) Enqueue(data ...T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.data.Enqueue(data...)
}

func (q *ThreadsafeQueue[T]) Dequeue() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.data.Dequeue()
}

func (q *ThreadsafeQueue[T]) Peek() T {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.data.Peek()
}

func (q *ThreadsafeQueue[T]) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.data.IsEmpty()
}

func (q *ThreadsafeQueue[T]) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.data.Size()
}

func (q *ThreadsafeQueue[T]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.data.Clear()
}

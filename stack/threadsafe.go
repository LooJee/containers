package stack

import "sync"

// ThreadsafeStack is a thread-safe stack.
type ThreadsafeStack[T comparable] struct {
	lock sync.RWMutex
	data *Stack[T]
}

// BuildThreadsafeStack returns a new ThreadsafeStack.
func BuildThreadsafeStack[T comparable](data ...T) *ThreadsafeStack[T] {
	return &ThreadsafeStack[T]{
		data: BuildStack[T](data...),
		lock: sync.RWMutex{},
	}
}

// Push adds an element to the top of the stack.
func (s *ThreadsafeStack[T]) Push(data ...T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data.Push(data...)
}

// Pop returns the top element of the stack and removes it.
func (s *ThreadsafeStack[T]) Pop() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.data.Pop()
}

// Peek returns the top element of the stack without removing it.
func (s *ThreadsafeStack[T]) Peek() T {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.data.Peek()
}

// IsEmpty returns true if the stack is empty.
func (s *ThreadsafeStack[T]) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.data.IsEmpty()
}

// Size returns the number of elements in the stack.
func (s *ThreadsafeStack[T]) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.data.Size()
}

// Clear removes all elements from the stack.
func (s *ThreadsafeStack[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data.Clear()
}

package set

import (
	"sync"
)

type ThreadsafeSet[T comparable] struct {
	lock sync.RWMutex
	set  *Set[T]
}

func BuildThreadsafeSet[T comparable](data ...T) *ThreadsafeSet[T] {
	return &ThreadsafeSet[T]{
		lock: sync.RWMutex{},
		set:  BuildSet[T](data...),
	}
}

func (s *ThreadsafeSet[T]) Insert(datas ...T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.set.Insert(datas...)
}

func (s *ThreadsafeSet[T]) Del(item T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.set.Del(item)
}

func (s *ThreadsafeSet[T]) Contains(item T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.set.Contains(item)
}

func (s *ThreadsafeSet[T]) Size() int {
	return s.set.Size()
}

func (s *ThreadsafeSet[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.set.Clear()
}

func (s *ThreadsafeSet[T]) Clone() *ThreadsafeSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return BuildThreadsafeSet(s.set.ToSlice()...)
}

func (s *ThreadsafeSet[T]) Diff(other *ThreadsafeSet[T]) *ThreadsafeSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	return BuildThreadsafeSet[T](s.set.Diff(other.set).ToSlice()...)
}

func (s *ThreadsafeSet[T]) Union(other *ThreadsafeSet[T]) *ThreadsafeSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	return BuildThreadsafeSet[T](s.set.Union(other.set).ToSlice()...)
}

func (s *ThreadsafeSet[T]) Range(fn func(T)) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.set.Range(fn)
}

func (s *ThreadsafeSet[T]) Intersect(other *ThreadsafeSet[T]) *ThreadsafeSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	return BuildThreadsafeSet[T](s.set.Intersect(other.set).ToSlice()...)
}

func (s *ThreadsafeSet[T]) Equal(other *ThreadsafeSet[T]) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	return s.set.Equal(other.set)
}

func (s *ThreadsafeSet[T]) IsEmpty() bool {
	return s.set.IsEmpty()
}

package set

import (
	"sync"
)

type Threadsafe[T comparable] struct {
	lock sync.RWMutex
	set  *ThreadUnsafe[T]
}

func BuildThreadSafe[T comparable](data ...T) *Threadsafe[T] {
	return &Threadsafe[T]{
		lock: sync.RWMutex{},
		set:  BuildThreadUnsafe[T](data...),
	}
}

func (s *Threadsafe[T]) Insert(datas ...T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.set.Insert(datas...)
}

func (s *Threadsafe[T]) Del(item T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.set.Del(item)
}

func (s *Threadsafe[T]) Contains(item T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.set.Contains(item)
}

func (s *Threadsafe[T]) Size() int {
	return s.set.Size()
}

func (s *Threadsafe[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.set.Clear()
}

func (s *Threadsafe[T]) Clone() *Threadsafe[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return BuildThreadSafe(s.set.ToSlice()...)
}

func (s *Threadsafe[T]) Diff(other *Threadsafe[T]) *Threadsafe[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	return BuildThreadSafe[T](s.set.Diff(other.set).ToSlice()...)
}

func (s *Threadsafe[T]) Union(other *Threadsafe[T]) *Threadsafe[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	return BuildThreadSafe[T](s.set.Union(other.set).ToSlice()...)
}

func (s *Threadsafe[T]) Range(fn func(T)) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.set.Range(fn)
}

func (s *Threadsafe[T]) Intersect(other *Threadsafe[T]) *Threadsafe[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	return BuildThreadSafe[T](s.set.Intersect(other.set).ToSlice()...)
}

func (s *Threadsafe[T]) Equal(other *Threadsafe[T]) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	return s.set.Equal(other.set)
}

func (s *Threadsafe[T]) IsEmpty() bool {
	return s.set.IsEmpty()
}

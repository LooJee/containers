package set

import (
	"sync"
)

type Threadsafe struct {
	lock sync.RWMutex
	set  *ThreadUnsafe
}

func BuildThreadSafe(data ...interface{}) (*Threadsafe, error) {
	s := &Threadsafe{
		lock: sync.RWMutex{},
	}

	var err error
	s.set, err = BuildThreadUnsafe(data...)

	return s, err
}

func (s *Threadsafe) Insert(datas ...interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.set.Insert(datas...)
}

func (s *Threadsafe) Del(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.set.Del(item)
}

func (s *Threadsafe) Contains(item interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.set.Contains(item)
}

func (s *Threadsafe) Size() int {
	return s.set.Size()
}

func (s *Threadsafe) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.set.Clear()
}

func (s *Threadsafe) Clone() *Threadsafe {
	s.lock.RLock()
	defer s.lock.RUnlock()

	cloner, _ := BuildThreadSafe(s.set.ToSlice()...)

	return cloner
}

func (s *Threadsafe) Diff(other *Threadsafe) (*Threadsafe, error) {
	var err error
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	diffSet, _ := BuildThreadSafe()
	diffSet.set, err = s.set.Diff(other.set)

	return diffSet, err
}

func (s *Threadsafe) Union(other *Threadsafe) (*Threadsafe, error) {
	var err error

	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	unionSet, _ := BuildThreadSafe()
	unionSet.set, err = s.set.Union(other.set)
	return unionSet, err
}

func (s *Threadsafe) Range(fn func(interface{})) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.set.Range(fn)
}

func (s *Threadsafe) Intersect(other *Threadsafe) (*Threadsafe, error) {
	var err error 
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	intersectionSet, _ := BuildThreadSafe()
	intersectionSet.set, err = s.set.Intersect(other.set)

	return intersectionSet, err
}

func (s *Threadsafe) Equal(other *Threadsafe) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	other.lock.RLock()
	defer other.lock.RUnlock()

	return s.set.Equal(other.set)
}

func (s *Threadsafe) IsEmpty() bool {
	return s.set.IsEmpty()
}

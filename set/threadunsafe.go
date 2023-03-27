package set

type ThreadUnsafe[T comparable] struct {
	set map[T]struct{}
}

func BuildThreadUnsafe[T comparable](data ...T) *ThreadUnsafe[T] {
	s := &ThreadUnsafe[T]{set: make(map[T]struct{})}

	s.Insert(data...)

	return s
}

func (s *ThreadUnsafe[T]) Insert(datas ...T) {
	for _, data := range datas {
		s.set[data] = struct{}{}
	}
}

func (s *ThreadUnsafe[T]) Contains(data T) bool {
	_, ok := s.set[data]
	return ok
}

func (s *ThreadUnsafe[T]) Del(data T) {
	delete(s.set, data)
}

func (s *ThreadUnsafe[T]) Range(fn func(data T)) {
	for data := range s.set {
		fn(data)
	}
}

func (s *ThreadUnsafe[T]) Size() int {
	return len(s.set)
}

func (s *ThreadUnsafe[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *ThreadUnsafe[T]) Clone() *ThreadUnsafe[T] {
	other := &ThreadUnsafe[T]{
		set: make(map[T]struct{}),
	}

	for data := range s.set {
		other.set[data] = struct{}{}
	}

	return other
}

// Equal returns true if s and other are equal.
func (s *ThreadUnsafe[T]) Equal(other *ThreadUnsafe[T]) bool {
	if s.Size() != other.Size() {
		return false
	}

	for data := range s.set {
		if !other.Contains(data) {
			return false
		}
	}

	return true
}

// Clear clears the set.
func (s *ThreadUnsafe[T]) Clear() {
	s.set = map[T]struct{}{}
}

func (s *ThreadUnsafe[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.set))

	for data := range s.set {
		slice = append(slice, data)
	}

	return slice
}

// Union returns a new set which contains the union of s and other.
func (s *ThreadUnsafe[T]) Union(other *ThreadUnsafe[T]) *ThreadUnsafe[T] {
	union := s.Clone()

	other.Range(func(data T) {
		union.Insert(data)
	})

	return union
}

// Diff returns a new set which contains the difference between s and other.
func (s *ThreadUnsafe[T]) Diff(other *ThreadUnsafe[T]) *ThreadUnsafe[T] {
	differ := BuildThreadUnsafe[T]()

	s.Range(func(data T) {
		if !other.Contains(data) {
			differ.Insert(data)
		}
	})

	return differ
}

// Intersect return a new set which contains the intersection between s and other.
func (s *ThreadUnsafe[T]) Intersect(other *ThreadUnsafe[T]) *ThreadUnsafe[T] {
	intersect := BuildThreadUnsafe[T]()

	s.Range(func(data T) {
		if other.Contains(data) {
			intersect.Insert(data)
		}
	})

	return intersect
}

package set

type Set[T comparable] struct {
	set map[T]struct{}
}

func BuildSet[T comparable](data ...T) *Set[T] {
	s := &Set[T]{set: make(map[T]struct{})}

	s.Insert(data...)

	return s
}

func (s *Set[T]) Insert(datas ...T) {
	for _, data := range datas {
		s.set[data] = struct{}{}
	}
}

func (s *Set[T]) Contains(data T) bool {
	_, ok := s.set[data]
	return ok
}

func (s *Set[T]) Del(data T) {
	delete(s.set, data)
}

func (s *Set[T]) Range(fn func(data T)) {
	for data := range s.set {
		fn(data)
	}
}

func (s *Set[T]) Size() int {
	return len(s.set)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set[T]) Clone() *Set[T] {
	other := &Set[T]{
		set: make(map[T]struct{}),
	}

	for data := range s.set {
		other.set[data] = struct{}{}
	}

	return other
}

// Equal returns true if s and other are equal.
func (s *Set[T]) Equal(other *Set[T]) bool {
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
func (s *Set[T]) Clear() {
	s.set = map[T]struct{}{}
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.set))

	for data := range s.set {
		slice = append(slice, data)
	}

	return slice
}

// Union returns a new set which contains the union of s and other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := s.Clone()

	other.Range(func(data T) {
		union.Insert(data)
	})

	return union
}

// Diff returns a new set which contains the difference between s and other.
func (s *Set[T]) Diff(other *Set[T]) *Set[T] {
	differ := BuildSet[T]()

	s.Range(func(data T) {
		if !other.Contains(data) {
			differ.Insert(data)
		}
	})

	return differ
}

// Intersect return a new set which contains the intersection between s and other.
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	intersect := BuildSet[T]()

	s.Range(func(data T) {
		if other.Contains(data) {
			intersect.Insert(data)
		}
	})

	return intersect
}

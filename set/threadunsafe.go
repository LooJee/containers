package set

import (
	"reflect"
)

type ThreadUnsafe struct {
	dataKind reflect.Kind
	set      map[interface{}]struct{}
}

var _ *ThreadUnsafe = (*ThreadUnsafe)(nil)

func BuildThreadUnsafe(data ...interface{}) (*ThreadUnsafe, error) {
	s := &ThreadUnsafe{set: make(map[interface{}]struct{})}

	err := s.Insert(data...)

	return s, err
}

func (s *ThreadUnsafe) isValidDataKind(dataKind reflect.Kind) bool {
	_, ok := validKind[dataKind]
	return ok
}

func (s *ThreadUnsafe) isSuitedDataKind(dataKind reflect.Kind) bool {
	return s.dataKind == reflect.Invalid || s.dataKind == dataKind
}

func (s *ThreadUnsafe) Insert(datas ...interface{}) error {
	for _, data := range datas {
		dataKind := reflect.TypeOf(data).Kind()

		if !s.isValidDataKind(dataKind) {
			return &InvalidDataTypeErr{DataType: dataKind.String()}
		}

		if !s.isSuitedDataKind(dataKind) {
			return &UnsuitableTypeErr{Want: s.dataKind.String(), Got: dataKind.String()}
		}

		s.set[data] = struct{}{}
		s.dataKind = dataKind
	}

	return nil
}

func (s *ThreadUnsafe) Contains(data interface{}) bool {
	_, ok := s.set[data]
	return ok
}

func (s *ThreadUnsafe) Del(data interface{}) {
	delete(s.set, data)
}

func (s *ThreadUnsafe) Range(fn func(data interface{})) {
	for data := range s.set {
		fn(data)
	}
}

func (s *ThreadUnsafe) Size() int {
	return len(s.set)
}

func (s *ThreadUnsafe) IsEmpty() bool {
	return s.Size() == 0
}

func (s *ThreadUnsafe) Clone() *ThreadUnsafe {
	other := &ThreadUnsafe{
		dataKind: s.dataKind,
		set:      make(map[interface{}]struct{}),
	}

	for data := range s.set {
		other.set[data] = struct{}{}
	}

	return other
}

// Equal returns true if s and other are equal.
func (s *ThreadUnsafe) Equal(other *ThreadUnsafe) bool {
	if _, err := s.assert(other); err != nil {
		return false
	}

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
func (s *ThreadUnsafe) Clear() {
	if s.dataKind == reflect.Invalid {
		return
	}

	s.dataKind = reflect.Invalid
	s.set = map[interface{}]struct{}{}
}

func (s *ThreadUnsafe) ToSlice() []interface{} {
	if s.dataKind == reflect.Invalid {
		return nil
	}

	slice := make([]interface{}, 0, len(s.set))

	for data := range s.set {
		slice = append(slice, data)
	}

	return slice
}

// Union returns a new set which contains the union of s and other.
func (s *ThreadUnsafe) Union(other *ThreadUnsafe) (*ThreadUnsafe, error) {
	if _, err := s.assert(other); err != nil {
		return nil, err
	}

	union := s.Clone()

	other.Range(func(data interface{}) {
		union.Insert(data)
	})

	return union, nil
}

// Diff returns a new set which contains the difference between s and other.
func (s *ThreadUnsafe) Diff(other *ThreadUnsafe) (*ThreadUnsafe, error) {
	if _, err := s.assert(other); err != nil {
		return nil, err
	}

	differ, _ := BuildThreadUnsafe()

	s.Range(func(data interface{}) {
		if !other.Contains(data) {
			differ.Insert(data)
		}
	})

	return differ, nil
}

// Intersect return a new set which contains the intersection between s and other.
func (s *ThreadUnsafe) Intersect(other *ThreadUnsafe) (*ThreadUnsafe, error) {
	if _, err := s.assert(other); err != nil {
		return nil, err
	}

	intersect, _ := BuildThreadUnsafe()

	s.Range(func(data interface{}) {
		if other.Contains(data) {
			intersect.Insert(data)
		}
	})

	return intersect, nil
}

func (s *ThreadUnsafe) assert(other *ThreadUnsafe) (*ThreadUnsafe, error) {
	if s.dataKind != other.dataKind {
		return nil, &UnsuitableTypeErr{Want: s.dataKind.String(), Got: other.dataKind.String()}
	}

	return other, nil
}

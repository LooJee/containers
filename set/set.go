package set

import (
	"reflect"
)

var (
	validKind = make(map[reflect.Kind]struct{})
)

type Set struct {
	dataKind reflect.Kind
	set      map[interface{}]struct{}
}

func init() {
	for i := reflect.Int; i <= reflect.Float64; i++ {
		validKind[i] = struct{}{}
	}

	validKind[reflect.String] = struct{}{}
}

func Build(data ...interface{}) (*Set, error) {
	s := &Set{set: make(map[interface{}]struct{})}

	err := s.Insert(data...)

	return s, err
}

func (s *Set) isValidDataKind(dataKind reflect.Kind) bool {
	_, ok := validKind[dataKind]
	return ok
}

func (s *Set) isSuitedDataKind(dataKind reflect.Kind) bool {
	return s.dataKind == reflect.Invalid || s.dataKind == dataKind
}

func (s *Set) Insert(datas ...interface{}) error {
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

func (s *Set) Contains(data interface{}) bool {
	dataKind := reflect.TypeOf(data).Kind()

	if !s.isValidDataKind(dataKind) || !s.isSuitedDataKind(dataKind) {
		return false
	}

	_, ok := s.set[data]
	return ok
}

func (s *Set) Del(data interface{}) error {
	dataKind := reflect.TypeOf(data).Kind()

	if !s.isValidDataKind(dataKind) {
		return &InvalidDataTypeErr{DataType: dataKind.String()}
	}

	if !s.isSuitedDataKind(dataKind) {
		return &UnsuitableTypeErr{Want: s.dataKind.String(), Got: dataKind.String()}
	}

	delete(s.set, data)

	return nil
}

func (s *Set) Range(fn func(data interface{})) {
	for data := range s.set {
		fn(data)
	}
}

func (s *Set) Size() int {
	return len(s.set)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) Clone() *Set {
	other := &Set{
		dataKind: s.dataKind,
		set:      make(map[interface{}]struct{}),
	}

	for data := range s.set {
		other.set[data] = struct{}{}
	}

	return other
}

func (s *Set) Equal(other *Set) bool {
	if s.dataKind != other.dataKind {
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

func (s *Set) Clear() {
	if s.dataKind == reflect.Invalid {
		return
	}

	s.dataKind = reflect.Invalid
	s.set = map[interface{}]struct{}{}
}

func (s *Set) ToSlice() []interface{} {
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
func (s *Set) Union(other *Set) (*Set, error) {
	if s.dataKind != other.dataKind {
		return nil, &UnsuitableTypeErr{Want: s.dataKind.String(), Got: other.dataKind.String()}
	}

	union := s.Clone()

	other.Range(func(data interface{}) {
		union.set[data] = struct{}{}
	})

	return union, nil
}

// Diff returns a new set which contains the difference between s and other.
func (s *Set) Diff(other *Set) (*Set, error) {
	if s.dataKind != other.dataKind {
		return nil, &UnsuitableTypeErr{Want: s.dataKind.String(), Got: other.dataKind.String()}
	}

	differ, _ := Build()

	s.Range(func(data interface{}) {
		if !other.Contains(data) {
			differ.Insert(data)
		}
	})

	return differ, nil
}

// return a new set which contains the intersection between s and other.
func (s *Set) Intersect(other *Set) (*Set, error) {
	if s.dataKind != other.dataKind {
		return nil, &UnsuitableTypeErr{Want: s.dataKind.String(), Got: other.dataKind.String()}
	}

	intersect, _ := Build()

	s.Range(func(data interface{}) {
		if other.Contains(data) {
			intersect.Insert(data)
		}
	})

	return intersect, nil
}

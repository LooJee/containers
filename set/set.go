package set

import "reflect"

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

func NewSet() Set {
	return Set{set: make(map[interface{}]struct{})}
}

func (s *Set) isValidDataKind(dataKind reflect.Kind) bool {
	_, ok := validKind[dataKind]
	return ok
}

func (s *Set) isSuitedDataKind(dataKind reflect.Kind) bool {
	return s.dataKind == reflect.Invalid || s.dataKind == dataKind
}

func (s *Set) Insert(data interface{}) error {
	dataKind := reflect.TypeOf(data).Kind()

	if !s.isValidDataKind(dataKind) {
		return &InvalidDataTypeErr{DataType: dataKind.String()}
	}

	if !s.isSuitedDataKind(dataKind) {
		return &UnsuitableTypeErr{Want: s.dataKind.String(), Got: dataKind.String()}
	}

	s.set[data] = struct{}{}
	s.dataKind = dataKind

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

func (s *Set) Iter(fn func(data interface{})) {
	cloner := s.Clone()

	for data := range cloner.set {
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

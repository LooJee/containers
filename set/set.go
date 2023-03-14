package set

import (
	"reflect"
	"sync"

	"github.com/loojee/containers/locker"
)

var (
	validKind = make(map[reflect.Kind]struct{})
)

type Set struct {
	mu       locker.Locker
	dataKind reflect.Kind
	set      map[interface{}]struct{}
}

func init() {
	for i := reflect.Int; i <= reflect.Float64; i++ {
		validKind[i] = struct{}{}
	}

	validKind[reflect.String] = struct{}{}
}

func New(threadSafe bool) Set {
	return Set{set: make(map[interface{}]struct{}), mu: &locker.NoobLocker{}}
}

func NewThreadSafe() Set {
	return Set{set: make(map[interface{}]struct{}), mu: &sync.RWMutex{}}
}

func (s *Set) isValidDataKind(dataKind reflect.Kind) bool {
	_, ok := validKind[dataKind]
	return ok
}

func (s *Set) isSuitedDataKind(dataKind reflect.Kind) bool {
	return s.dataKind == reflect.Invalid || s.dataKind == dataKind
}

func (s *Set) Insert(datas ...interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

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
	s.mu.RLock()
	defer s.mu.RUnlock()

	dataKind := reflect.TypeOf(data).Kind()

	if !s.isValidDataKind(dataKind) || !s.isSuitedDataKind(dataKind) {
		return false
	}

	_, ok := s.set[data]
	return ok
}

func (s *Set) Del(data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

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
	s.mu.Lock()
	defer s.mu.Unlock()
	other := &Set{
		dataKind: s.dataKind,
		set:      make(map[interface{}]struct{}),
	}

	other.mu = reflect.New(reflect.TypeOf(s.mu).Elem()).Interface().(locker.Locker)

	for data := range s.set {
		other.set[data] = struct{}{}
	}

	return other
}

func (s *Set) Equal(other *Set) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

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
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.dataKind == reflect.Invalid {
		return
	}

	s.dataKind = reflect.Invalid
	s.set = map[interface{}]struct{}{}
}

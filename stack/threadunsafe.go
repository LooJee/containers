package stack

// Stack is a thread-unsafe stack.
type Stack[T comparable] struct {
	data []T
}

// BuildStack returns a new Stack.
func BuildStack[T comparable](data ...T) *Stack[T] {
	s := &Stack[T]{data: make([]T, 0)}

	s.Push(data...)

	return s
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(data ...T) {
	s.data = append(s.data, data...)
}

// Pop returns the top element of the stack and removes it.
func (s *Stack[T]) Pop() T {
	var data T
	if s.IsEmpty() {
		return data
	}

	data = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return data
}

// Peek returns the top element of the stack without removing it.
func (s *Stack[T]) Peek() T {
	var data T
	if s.IsEmpty() {
		return data
	}

	return s.data[len(s.data)-1]
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.data = []T{}
}

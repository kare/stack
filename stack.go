package stack

import "errors"

// ErrEmptyStack error signals that the stack is empty.
var ErrEmptyStack = errors.New("stack is empty")

// Stack is a stack data structure.
type Stack[T any] []T

// New creates an empty string stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push a value to the stack.
func (s *Stack[T]) Push(value T) {
	*s = append([]T{value}, *s...)
}

// Pop removes the most recently added item.
// If called on an empty stack will return ErrEmptyStack error.
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(*s) == 0 {
		return zero, ErrEmptyStack
	}
	value := (*s)[0]
	(*s)[0] = zero
	(*s) = (*s)[1:]
	return value, nil
}

// Slice returns the stack contents as a slice of string's.
func (s *Stack[T]) Slice() []T {
	return []T(*s)
}

// Peek returns the most recently added item.
// If called on an empty stack will return ErrEmptyStack error.
func (s *Stack[T]) Peek() (T, error) {
	if len(*s) == 0 {
		var zero T
		return zero, ErrEmptyStack
	}
	return (*s)[0], nil
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[_]) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of items in the stack.
func (s *Stack[T]) Len() int {
	return len(*s)
}

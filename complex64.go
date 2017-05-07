// Code generated by "gends"; DO NOT EDIT.

package stack // import "kkn.fi/stack"

// Complex64Stack is a stack data structure.
type Complex64Stack []complex64

// NewComplex64 creates an empty complex64 stack.
func NewComplex64() *Complex64Stack {
	return &Complex64Stack{}
}

// Push a value to stack
func (s *Complex64Stack) Push(value complex64) {
	(*s) = append([]complex64{value}, (*s)...)
}

// Pop removes the most recently added item
func (s *Complex64Stack) Pop() complex64 {
	value := (*s)[0]
	(*s) = (*s)[1:]
	return value
}

// Slice returns the stack contents as an complex64 slice.
func (s *Complex64Stack) Slice() []complex64 {
	return []complex64(*s)
}

// Peek returns the most recently added item.
func (s *Complex64Stack) Peek() complex64 {
	return (*s)[0]
}

// IsEmpty returns true if the stack is empty.
func (s *Complex64Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of items in the stack.
func (s *Complex64Stack) Len() int {
	return len(*s)
}

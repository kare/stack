// Code generated by "gends"; DO NOT EDIT.

package stack // import "kkn.fi/stack"

// Complex128Stack is a stack data structure.
type Complex128Stack []complex128

// NewComplex128 creates an empty complex128 stack.
func NewComplex128() *Complex128Stack {
	return &Complex128Stack{}
}

// Push a value to stack
func (s *Complex128Stack) Push(value complex128) {
	(*s) = append([]complex128{value}, (*s)...)
}

// Pop removes the most recently added item
func (s *Complex128Stack) Pop() complex128 {
	value := (*s)[0]
	(*s) = (*s)[1:]
	return value
}

// Slice returns the stack contents as an complex128 slice.
func (s *Complex128Stack) Slice() []complex128 {
	return []complex128(*s)
}

// Peek returns the most recently added item.
func (s *Complex128Stack) Peek() complex128 {
	return (*s)[0]
}

// IsEmpty returns true if the stack is empty.
func (s *Complex128Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of items in the stack.
func (s *Complex128Stack) Len() int {
	return len(*s)
}

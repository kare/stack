// Code generated by "gends"; DO NOT EDIT.

package stack // import "kkn.fi/stack"

// Complex64 is a stack data structure.
type Complex64 []complex64

// NewComplex64 creates an empty complex64 stack.
func NewComplex64() *Complex64 {
	return &Complex64{}
}

// Push a value to the stack.
func (s *Complex64) Push(value complex64) {
	(*s) = append([]complex64{value}, (*s)...)
}

// Pop removes the most recently added item.
// If called on an empty stack will return ErrEmptyStack error.
func (s *Complex64) Pop() (complex64, error) {
	if len(*s) == 0 {
		return 0, ErrEmptyStack
	}
	value := (*s)[0]
	(*s) = (*s)[1:]
	return value, nil
}

// Slice returns the stack contents as a slice of complex64's.
func (s *Complex64) Slice() []complex64 {
	return []complex64(*s)
}

// Peek returns the most recently added item.
// If called on an empty stack will return ErrEmptyStack error.
func (s *Complex64) Peek() (complex64, error) {
	if len(*s) == 0 {
		return 0, ErrEmptyStack
	}
	return (*s)[0], nil
}

// IsEmpty returns true if the stack is empty.
func (s *Complex64) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of items in the stack.
func (s *Complex64) Len() int {
	return len(*s)
}

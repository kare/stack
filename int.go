// Code generated by "gends"; DO NOT EDIT.

package stack // import "kkn.fi/stack"

// Int is a stack data structure.
type Int []int

// NewInt creates an empty int stack.
func NewInt() *Int {
	return &Int{}
}

// Push a value to the stack.
func (s *Int) Push(value int) {
	(*s) = append([]int{value}, (*s)...)
}

// Pop removes the most recently added item.
// If called on an empty stack will return ErrEmptyStack error.
func (s *Int) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, ErrEmptyStack
	}
	value := (*s)[0]
	(*s) = (*s)[1:]
	return value, nil
}

// Slice returns the stack contents as a slice of int's.
func (s *Int) Slice() []int {
	return []int(*s)
}

// Peek returns the most recently added item.
// If called on an empty stack will return ErrEmptyStack error.
func (s *Int) Peek() (int, error) {
	if len(*s) == 0 {
		return 0, ErrEmptyStack
	}
	return (*s)[0], nil
}

// IsEmpty returns true if the stack is empty.
func (s *Int) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of items in the stack.
func (s *Int) Len() int {
	return len(*s)
}

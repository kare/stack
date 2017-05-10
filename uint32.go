// Code generated by "gends"; DO NOT EDIT.

package stack // import "kkn.fi/stack"

// Uint32 is a stack data structure.
type Uint32 []uint32

// NewUint32 creates an empty uint32 stack.
func NewUint32() *Uint32 {
	return &Uint32{}
}

// Push a value to the stack.
func (s *Uint32) Push(value uint32) {
	(*s) = append([]uint32{value}, (*s)...)
}

// Pop removes the most recently added item.
// If called on an empty stack will return ErrEmptyStack error.
func (s *Uint32) Pop() (uint32, error) {
	if len(*s) == 0 {
		return 0, ErrEmptyStack
	}
	value := (*s)[0]
	(*s) = (*s)[1:]
	return value, nil
}

// Slice returns the stack contents as a slice of uint32's.
func (s *Uint32) Slice() []uint32 {
	return []uint32(*s)
}

// Peek returns the most recently added item.
// If called on an empty stack will return ErrEmptyStack error.
func (s *Uint32) Peek() (uint32, error) {
	if len(*s) == 0 {
		return 0, ErrEmptyStack
	}
	return (*s)[0], nil
}

// IsEmpty returns true if the stack is empty.
func (s *Uint32) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of items in the stack.
func (s *Uint32) Len() int {
	return len(*s)
}

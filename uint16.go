// Code generated by "gends"; DO NOT EDIT.

package stack // import "kkn.fi/stack"

// Uint16Stack is a stack data structure.
type Uint16Stack []uint16

// NewUint16 creates an empty uint16 stack.
func NewUint16() *Uint16Stack {
	return &Uint16Stack{}
}

// Push a value to stack
func (s *Uint16Stack) Push(value uint16) {
	(*s) = append([]uint16{value}, (*s)...)
}

// Pop removes the most recently added item
func (s *Uint16Stack) Pop() uint16 {
	value := (*s)[0]
	(*s) = (*s)[1:]
	return value
}

// Slice returns the stack contents as an uint16 slice.
func (s *Uint16Stack) Slice() []uint16 {
	return []uint16(*s)
}

// Peek returns the most recently added item.
func (s *Uint16Stack) Peek() uint16 {
	return (*s)[0]
}

// IsEmpty returns true if the stack is empty.
func (s *Uint16Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of items in the stack.
func (s *Uint16Stack) Len() int {
	return len(*s)
}
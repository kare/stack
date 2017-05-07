// Code generated by "gends"; DO NOT EDIT.

package stack // import "kkn.fi/stack"

// RuneStack is a stack data structure.
type RuneStack []rune

// NewRune creates an empty rune stack.
func NewRune() *RuneStack {
	return &RuneStack{}
}

// Push a value to stack
func (s *RuneStack) Push(value rune) {
	(*s) = append([]rune{value}, (*s)...)
}

// Pop removes the most recently added item
func (s *RuneStack) Pop() rune {
	value := (*s)[0]
	(*s) = (*s)[1:]
	return value
}

// Slice returns the stack contents as an rune slice.
func (s *RuneStack) Slice() []rune {
	return []rune(*s)
}

// Peek returns the most recently added item.
func (s *RuneStack) Peek() rune {
	return (*s)[0]
}

// IsEmpty returns true if the stack is empty.
func (s *RuneStack) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of items in the stack.
func (s *RuneStack) Len() int {
	return len(*s)
}

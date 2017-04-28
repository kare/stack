package stack // import "kkn.fi/stack"

// IntStack is a stack data structure.
type IntStack []int

// NewInt creates an empty int stack.
func NewInt() *IntStack {
	return &IntStack{}
}

// Push a value to stack
func (s *IntStack) Push(value int) {
	(*s) = append([]int{value}, (*s)...)
}

// Pop removes the most recently added item
func (s *IntStack) Pop() int {
	value := (*s)[0]
	(*s) = (*s)[1:]
	return value
}

// Slice returns the stack contents as an int slice.
func (s *IntStack) Slice() []int {
	return []int(*s)
}

// Peek returns the most recently added item.
func (s *IntStack) Peek() int {
	return (*s)[0]
}

// IsEmpty returns true if the stack is empty.
func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of items in the stack.
func (s *IntStack) Len() int {
	return len(*s)
}
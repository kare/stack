package stack

import "testing"

func TestIntStackSimple(t *testing.T) {
	stack := NewInt()
	stack.Push(1)
	l := stack.Len()
	if l != 1 {
		t.Errorf("Stack Len() should be 1, but was %v", l)
	}
	value := stack.Pop()
	expected := 1
	if value != expected {
		t.Errorf("Stack should have returned %v, but returned %v", expected, value)
	}
	l = stack.Len()
	if l != 0 {
		t.Errorf("Stack Len() should be 0, but was %v", l)
	}
}

func TestIntStackFiveValues(t *testing.T) {
	stack := NewInt()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	for i := 4; i >= 0; i-- {
		value := stack.Pop()
		expected := i
		if value != expected {
			t.Errorf("Stack should have returned %v, but returned %v", expected, value)
		}
	}
}

func TestIntStackIsEmpty(t *testing.T) {
	stack := NewInt()
	if !stack.IsEmpty() {
		t.Errorf("Expected an empty stack, but got: %v", stack)
	}
	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("Did not expect an empty stack, but it was empty")
	}
}

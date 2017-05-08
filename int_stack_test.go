package stack

import "testing"

func TestIntStackSimple(t *testing.T) {
	stack := NewInt()
	stack.Push(1)
	l := stack.Len()
	if l != 1 {
		t.Errorf("Stack Len() should be 1, but was %v", l)
	}
	value, _ := stack.Pop()
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
		value, _ := stack.Pop()
		expected := i
		if value != expected {
			t.Errorf("Stack should have returned %v, but returned %v", expected, value)
		}
	}
}

func TestIntStackPopEmpty(t *testing.T) {
	stack := NewInt()
	zero, err := stack.Pop()
	if err.Error() != "stack is empty" {
		t.Errorf("expected peek to return an error, but got %d, %v", zero, err)
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

func TestIntStackPeek(t *testing.T) {
	stack := NewInt()
	stack.Push(1)
	one, _ := stack.Peek()
	if one != 1 {
		t.Errorf("expected peek to return 1, but got %d", one)
	}
}

func TestIntStackPeekEmpty(t *testing.T) {
	stack := NewInt()
	zero, err := stack.Peek()
	if err.Error() != "stack is empty" {
		t.Errorf("expected peek to return an error, but got %d, %v", zero, err)
	}
}

func TestIntStackSlice(t *testing.T) {
	stack := NewInt()
	stack.Push(1)
	stack.Push(2)
	slice := stack.Slice()
	if slice[0] != 2 {
		t.Errorf("expected slice[0] to equal %d", 2)
	}
	if slice[1] != 1 {
		t.Errorf("expected slice[1] to equal %d", 1)
	}
}

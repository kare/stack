package stack

import "testing"

func TestIntStackSimple(t *testing.T) {
	s := New[int]()
	s.Push(1)
	l := s.Len()
	if l != 1 {
		t.Errorf("Stack Len() should be 1, but was %v", l)
	}
	value, _ := s.Pop()
	expected := 1
	if value != expected {
		t.Errorf("Stack should have returned %v, but returned %v", expected, value)
	}
	l = s.Len()
	if l != 0 {
		t.Errorf("Stack Len() should be 0, but was %v", l)
	}
}

func TestIntStackFiveValues(t *testing.T) {
	s := New[int]()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	for i := 4; i >= 0; i-- {
		value, _ := s.Pop()
		expected := i
		if value != expected {
			t.Errorf("Stack should have returned %v, but returned %v", expected, value)
		}
	}
}

func TestIntStackPopEmpty(t *testing.T) {
	s := New[int]()
	zero, err := s.Pop()
	if err.Error() != "stack is empty" {
		t.Errorf("expected peek to return an error, but got %d, %v", zero, err)
	}
}

func TestIntStackIsEmpty(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Errorf("Expected an empty stack, but got: %v", s)
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Error("Did not expect an empty stack, but it was empty")
	}
}

func TestIntStackPeek(t *testing.T) {
	s := New[int]()
	s.Push(1)
	one, _ := s.Peek()
	if one != 1 {
		t.Errorf("expected peek to return 1, but got %d", one)
	}
}

func TestIntStackPeekEmpty(t *testing.T) {
	s := New[int]()
	zero, err := s.Peek()
	if err.Error() != "stack is empty" {
		t.Errorf("expected peek to return an error, but got %d, %v", zero, err)
	}
}

func TestIntStackSlice(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	slice := s.Slice()
	if slice[0] != 2 {
		t.Errorf("expected slice[0] to equal %d", 2)
	}
	if slice[1] != 1 {
		t.Errorf("expected slice[1] to equal %d", 1)
	}
}

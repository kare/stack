package stack // import "kkn.fi/stack"

import "errors"

//go:generate gends templates.json stack.tmpl

// ErrEmptyStack error signals that the stack is empty.
var ErrEmptyStack = errors.New("stack is empty")

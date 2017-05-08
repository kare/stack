package stack // import "kkn.fi/stack"

import "errors"

//go:generate gends templates.json stack.tmpl

var ErrEmptyStack = errors.New("stack is empty")

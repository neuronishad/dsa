package ds

import (
	"dsa/err"
)

const (
	ErrStackFull  rr.Err  = "stack full"
	ErrStackEmpty err.Err = "stack empty"
)

type Stack struct {
	stack  []string
	length int
	top    int
}

func (s *Stack) Push(data string) error {
	if s.length != -1 {
		if s.top+1 == s.length {
			return ErrStackFull
		}

		s.stack[s.top+1] = data
		s.top++
	} else {
		s.stack = append(s.stack, data)
		s.top++
	}

	return nil
}

func (s *Stack) Peek() (string, error) {
	if s.top == -1 {
		return "", ErrStackEmpty
	}

	return s.stack[s.top], nil
}

func NewFixedStack(length int) *Stack {
	return &Stack{
		stack:  make([]string, length),
		length: length,
		top:    -1,
	}
}

func NewDynamicStack() *Stack {
	return &Stack{
		stack:  []string{},
		length: -1,
		top:    -1,
	}
}

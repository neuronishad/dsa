package ds

import (
	"dsa/err"

	"golang.org/x/exp/constraints"
)

const (
	ErrStackFull  err.Err = "stack full"
	ErrStackEmpty err.Err = "stack empty"
)

type Stack[T any, M constraints.Signed] struct {
	stack  []T
	length M
	top    M
}

func (s *Stack[T, M]) Push(data T) error {
	if s.length != M(-1) {
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

func (s *Stack[T, M]) Pop() (T, error) {
	if s.top == M(-1) {
		return *new(T), ErrStackEmpty
	}

	data := s.stack[s.top]
	if s.length != M(-1) {
		s.stack[s.top] = *new(T)
	} else {
		s.stack = s.stack[:s.top]
	}

	s.top--

	return data, nil
}

func (s *Stack[T, M]) Peek() (T, error) {
	if s.top == M(-1) {
		return *new(T), ErrStackEmpty
	}

	return s.stack[s.top], nil
}

func (s *Stack[T, M]) IsEmpty() bool {
	return s.top == M(-1)
}

func (s *Stack[T, M]) IsFull() bool {
	if s.length == M(-1) {
		return false
	} else {
		return s.top == s.length
	}
}

func NewFixedStack[T any, M constraints.Signed](length M) *Stack[T, M] {
	return &Stack[T, M]{
		stack:  make([]T, length),
		length: length,
		top:    M(-1),
	}
}

func NewDynamicStack[T any, M constraints.Signed]() *Stack[T, M] {
	return &Stack[T, M]{
		stack:  []T{},
		length: M(-1),
		top:    M(-1),
	}
}

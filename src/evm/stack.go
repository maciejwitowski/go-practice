package evm

import (
	"errors"
)

type Stack struct {
	items []Value
}

func NewStack(initialCapacity int) *Stack {
	return &Stack{
		items: make([]Value, initialCapacity),
	}
}

func (s *Stack) Push(item Value) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() (Value, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}

	lastIndex := len(s.items) - 1
	last := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return last, nil
}

func (s *Stack) peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

type Value interface {
}

type IntValue struct {
	val int
}

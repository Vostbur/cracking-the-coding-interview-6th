package chapter_3

import "errors"

type StackSlice struct {
	items []int
}

func (s *StackSlice) push(x int) {
	s.items = append(s.items, x)
}

func (s *StackSlice) pop() (int, error) {
	l := len(s.items) - 1
	if l == -1 {
		return -1, errors.New("Cannot pop. Stack is empty.")
	}
	ret := s.items[l]
	s.items = s.items[:l]
	return ret, nil
}

func (s StackSlice) peek() (int, error) {
	l := len(s.items) - 1
	if l == -1 {
		return -1, errors.New("Cannot peek. Stack is empty.")
	}
	return s.items[l], nil
}

func (s StackSlice) isEmpty() bool {
	return len(s.items) == 0
}

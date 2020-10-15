package chapter_3

import "errors"

type stackNode struct {
	value int
	prev  *stackNode
}

type Stack struct {
	top *stackNode
}

func (s *Stack) push(x int) {
	newNode := &stackNode{value: x}
	if s.top != nil {
		newNode.prev = s.top
	}
	s.top = newNode
}

func (s *Stack) pop() (int, error) {
	if s.top == nil {
		return -1, errors.New("Cannot pop. Stack is empty.")
	}
	x := s.top.value
	s.top = s.top.prev
	return x, nil
}

func (s Stack) peek() (int, error) {
	if s.top == nil {
		return -1, errors.New("Cannot peek. Stack is empty.")
	}
	return s.top.value, nil
}

func (s Stack) isEmpty() bool {
	return s.top == nil
}

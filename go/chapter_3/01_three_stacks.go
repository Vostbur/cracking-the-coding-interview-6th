package chapter_3

import "fmt"

type MultiStack struct {
	stacks    int
	stackSize int
	stack     []int
	sizes     []int
}

func GetMultiStack(stacks, stackSize int) *MultiStack {
	return &MultiStack{
		stacks,
		stackSize,
		make([]int, stacks*stackSize),
		make([]int, stacks),
	}
}

func (s *MultiStack) Push(stack, value int) error {
	if s.IsFull(stack) {
		return fmt.Errorf("Cannot push. Stack %d is full.", stack)
	}
	s.sizes[stack-1]++
	s.stack[s.Index(stack)] = value
	return nil
}

func (s *MultiStack) Pop(stack int) (int, error) {
	if s.IsEmpty(stack) {
		return -1, fmt.Errorf("Cannot pop. Stack %d is empty.", stack)
	}
	ind := s.Index(stack)
	ret := s.stack[ind]
	s.stack[ind] = 0
	s.sizes[stack-1]--
	return ret, nil
}

func (s MultiStack) Peek(stack int) (int, error) {
	if s.IsEmpty(stack) {
		return -1, fmt.Errorf("Cannot peek. Stack %d is empty.", stack)
	}
	return s.stack[s.Index(stack)], nil
}

func (s MultiStack) Index(stack int) int {
	stack--
	return stack*s.stackSize + s.sizes[stack] - 1
}

func (s MultiStack) IsEmpty(stack int) bool {
	return s.sizes[stack-1] == 0
}

func (s MultiStack) IsFull(stack int) bool {
	return s.sizes[stack-1] == s.stackSize
}

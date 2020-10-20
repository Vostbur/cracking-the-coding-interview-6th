// Как реализовать стек, в котором кроме стандартных функций push и pop будет
// поддерживаться функция min, возвращающая минимальный элемент? Все
// операции — push, pop и min — должны выполняться за время O(1).
package chapter_3

import "fmt"

type StackWithMin struct {
	stack     *Stack
	minValues *Stack
}

func GetStackWithMin() *StackWithMin {
	return &StackWithMin{GetStack(), GetStack()}
}

func (s *StackWithMin) Push(value int) {
	var m int
	if s.IsEmpty() {
		m = value
	} else {
		m, _ = s.Min()
	}

	if m > value {
		m = value
	}
	s.stack.push(value)
	s.minValues.push(m)
}

func (s *StackWithMin) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("Cannot pop. Stack is empty.")
	}

	ret, err := s.stack.pop()
	if err != nil {
		return -1, err
	}
	_, err = s.minValues.pop()
	if err != nil {
		return -1, err
	}

	return ret, nil
}

func (s StackWithMin) Peek() (int, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("Cannot peek. Stack is empty.")
	}
	return s.stack.peek()
}

func (s StackWithMin) Min() (int, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("Cannot get min element. Stack is empty.")
	}
	return s.minValues.peek()
}

func (s StackWithMin) IsEmpty() bool {
	return s.stack.isEmpty()
}

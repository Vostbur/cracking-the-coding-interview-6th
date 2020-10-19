package chapter_3

import "testing"

func TestMin(t *testing.T) {
	s := GetStackWithMin()
	s.Push(3)
	s.Push(2)
	s.Push(4)
	s.Push(1)
	if m, _ := s.Min(); m != 1 {
		t.Errorf("Expected: 1. Actual: %d", m)
	}
	_, _ = s.Pop()
	if m, _ := s.Min(); m != 2 {
		t.Errorf("Expected: 2. Actual: %d", m)
	}
	_, _ = s.Pop()
	_, _ = s.Pop()
	if m, _ := s.Min(); m != 3 {
		t.Errorf("Expected: 3. Actual: %d", m)
	}
}

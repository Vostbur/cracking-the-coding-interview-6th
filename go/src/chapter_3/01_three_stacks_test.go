package chapter_3

import (
	"reflect"
	"testing"
)

func TestThreeStacks(t *testing.T) {
	stack1 := []int{1, 2, 3, 4, 5}
	stack2 := []int{6, 7, 8}
	stack3 := []int{9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 9, 10, 0, 0, 0}

	s := GetMultiStack(3, 5)
	for _, i := range stack1 {
		_ = s.Push(1, i)
	}
	for _, i := range stack2 {
		_ = s.Push(2, i)
	}
	for _, i := range stack3 {
		_ = s.Push(3, i)
	}

	actual := []int{}
	actual = append(actual, s.stack...)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}

	pop1, err := s.Pop(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	pop2, err := s.Pop(2)
	if err != nil {
		t.Errorf(err.Error())
	}
	pop3, err := s.Pop(3)
	if err != nil {
		t.Errorf(err.Error())
	}

	expected = []int{1, 2, 3, 4, 0, 6, 7, 0, 0, 0, 9, 0, 0, 0, 0}
	actual = []int{}
	actual = append(actual, s.stack...)
	
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}

	poped := []int{pop1, pop2, pop3}
	expected = []int{5, 8, 10}
	if !reflect.DeepEqual(poped, expected) {
		t.Errorf("Poped: %v. Actual: %v", expected, poped)
	}
}

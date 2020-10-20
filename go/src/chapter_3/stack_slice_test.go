package chapter_3

import (
	"reflect"
	"sort"
	"testing"
)

func TestStackSlice(t *testing.T) {
	init := []int{1, 2, 3, 4, 5}
	s := &StackSlice{}
	for _, i := range init {
		s.push(i)
	}

	fifo := init[len(init)-1]
	if r, _ := s.peek(); r != fifo {
		t.Errorf("Expect: %d. Actual: %d\n", fifo, r)
	}

	actual := []int{}
	for {
		value, err := s.pop()
		if err != nil {
			break
		}
		actual = append(actual, value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(init)))
	if !reflect.DeepEqual(init, actual) {
		t.Errorf("Expect: %v. Actual: %v\n", init, actual)
	}

	if r := s.isEmpty(); r != true {
		t.Errorf("Expect: %v. Actual: %v\n", true, r)
	}

	if _, err := s.peek(); err == nil {
		t.Errorf("Stack is not empty.")
	}
}

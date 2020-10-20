package chapter_3

import (
	"reflect"
	"testing"
)

func TestSortedStack(t *testing.T) {
	init := []int{1, 3, 6, 5, 7, 2, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7}

	s := GetSortedStack()
	for _, i := range init {
		s.push(i)
	}

	actual := []int{}
	for {
		poped, err := s.pop()
		if err != nil {
			break
		}
		actual = append(actual, poped)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
}

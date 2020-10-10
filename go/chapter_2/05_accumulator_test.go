package chapter_2

import (
	"reflect"
	"testing"
)

func TestAccumulator(t *testing.T) {
	cases := []struct {
		l1, l2   []int
		expected []int
	}{
		{[]int{7, 1, 6}, []int{5, 9, 2}, []int{2, 1, 9}},
		{[]int{7, 1, 6, 1}, []int{5, 9, 2}, []int{2, 1, 9, 1}},
	}

	for _, c := range cases {
		ll1 := GetLinkedListFromSlice(c.l1)
		ll2 := GetLinkedListFromSlice(c.l2)
		if actual := Accumulator(*ll1, *ll2); !reflect.DeepEqual(actual.GetSlice(), c.expected) {
			t.Errorf("Expected: %v, actual: %v\n", c.expected, actual.GetSlice())
		}
	}
}

func TestAccumulatorInRow(t *testing.T) {
	cases := []struct {
		l1, l2   []int
		expected []int
	}{
		{[]int{6, 1, 7}, []int{2, 9, 5}, []int{9, 1, 2}},
		{[]int{1, 6, 1, 7}, []int{2, 9, 5}, []int{1, 9, 1, 2}},
	}

	for _, c := range cases {
		ll1 := GetLinkedListFromSlice(c.l1)
		ll2 := GetLinkedListFromSlice(c.l2)
		if actual := AccumulatorInRow(*ll1, *ll2); !reflect.DeepEqual(actual.GetSlice(), c.expected) {
			t.Errorf("Expected: %v, actual: %v\n", c.expected, actual.GetSlice())
		}
	}
}

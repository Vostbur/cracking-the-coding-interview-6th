package chapter_2

import "testing"

func LoopList(args []int, index int) *LinkedList {
	ll := GetLinkedListFromSlice(args)
	ll.tail.next, _ = ll.GetNodeByIndex(index)
	return ll
}

func TestLoopDetectFalse(t *testing.T) {
	cases := []struct {
		args []int
	}{
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1}},
		{[]int{}},
	}
	for _, c := range cases {
		ll := GetLinkedListFromSlice(c.args)
		if actual := ll.LoopDetect(); actual != nil {
			t.Errorf("Expected: nil, actual: %v\n", actual)
		}
	}
}

func TestLoopDetectTrue(t *testing.T) {
	cases := []struct {
		args  []int
		index int
	}{
		{[]int{0, 1, 2, 3, 4}, 0},
		{[]int{0, 1, 2, 3, 4}, 2},
		{[]int{0, 1, 2, 3, 4}, 4},
	}
	for _, c := range cases {
		ll := LoopList(c.args, c.index)
		if actual := ll.LoopDetect(); actual.value != c.index {
			t.Errorf("Expected: %v, actual: %v\n", c.index, actual)
		}
	}
}

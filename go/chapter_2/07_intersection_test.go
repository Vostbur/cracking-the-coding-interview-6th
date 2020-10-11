package chapter_2

import "testing"

func TestIntersectionTrue(t *testing.T) {
	l1 := GetLinkedListFromSlice([]int{1, 2, 3, 4, 5})
	l2 := l1
	if actual, _ := Intersection(l1, l2); !actual {
		t.Errorf("Expected: true, actual: %v\n", actual)
	}
}

func TestIntersectionFalse(t *testing.T) {
	cases := []struct {
		l1, l2   *LinkedList
		expected bool
	}{
		{
			GetLinkedListFromSlice([]int{1, 2, 3}),
			GetLinkedListFromSlice([]int{1, 2, 3}),
			false,
		},
		{
			GetLinkedListFromSlice([]int{}),
			GetLinkedListFromSlice([]int{}),
			false,
		},
	}
	for _, c := range cases {
		if actual, _ := Intersection(c.l1, c.l2); c.expected != actual {
			t.Errorf("Expected: %v, actual: %v\n", c.expected, actual)
		}
	}
}

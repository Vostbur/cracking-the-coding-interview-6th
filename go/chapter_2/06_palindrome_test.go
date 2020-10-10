package chapter_2

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		items    []int
		expected bool
	}{
		{[]int{1, 2, 2}, false},
		{[]int{1, 2, 1}, true},
		{[]int{2, 1, 1, 2}, true},
	}
	for _, c := range cases {
		ll := GetLinkedListFromSlice(c.items)
		actual := ll.IsPalindrome()
		if c.expected != actual {
			t.Fatalf("Expected: %v, actual: %v\n", c.expected, actual)
		}
	}
}

func TestIsPalindromeStack(t *testing.T) {
	cases := []struct {
		items    []int
		expected bool
	}{
		{[]int{1, 2, 2}, false},
		{[]int{1, 2, 1}, true},
		{[]int{2, 1, 1, 2}, true},
	}
	for _, c := range cases {
		ll := GetLinkedListFromSlice(c.items)
		actual := ll.IsPalindromeStack()
		if c.expected != actual {
			t.Fatalf("Expected: %v, actual: %v\n", c.expected, actual)
		}
	}
}

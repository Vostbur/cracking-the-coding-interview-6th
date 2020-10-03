package chapter_1

import "testing"

func TestModification(t *testing.T) {
	cases := []struct {
		s1 string
		s2 string
		expected bool
	}{
		{"abcd", "abcd", true},
		{"abcd", "abcc", true},
		{"abcd", "accc", false},
		{"abcd", "abcde", true},
		{"abcd", "abcdef", false},
		{"abcde", "abcd", true},
		{"abcdef", "abcd", false},
		{" ", "", true},
		{"", " ", true},
		{"", "", true},
	}
	for _, c := range cases {
		if actual := Modification(c.s1, c.s2); actual != c.expected {
			t.Errorf("Input: (%s, %s). Expected: %v. Actual: %v", c.s1, c.s2, c.expected, actual)
		}
	}
}
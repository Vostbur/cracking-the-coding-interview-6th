package chapter_1

import "testing"

func TestPermutation(t *testing.T) {
	cases := []struct {
		input    string
		output   string
		expected bool
	}{
		{"asd", "asd", false},
		{"asd", "asdf", false},
		{"", "", false},
		{"А вас", "Авас ", true},
		{"asd", "sda", true},
	}
	for _, c := range cases {
		if result := Permutation(c.input, c.output); result != c.expected {
			t.Errorf("Input: %s. Output: %s. Ecpected: %v. Result: %v\n", c.input, c.output, c.expected, result)
		}
	}
}

package chapter_1

import "testing"

func TestCompression(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"aaaabbbcdddd", "a4b3c1d4"},
		{"aaaaaaaaaaaaaaabcd", "a15b1c1d1"},
		{"abcd", "abcd"},
		{"ab", "ab"},
	}
	for _, c := range cases {
		if actual := Compression(c.input); actual != c.expected {
			t.Errorf("Input: %s. Expected: %s. Actual: %s\n", c.input, c.expected, actual)
		}
	}
	for _, c := range cases {
		if actual := CompressionWithSlice(c.input); actual != c.expected {
			t.Errorf("Input: %s. Expected: %s. Actual: %s\n", c.input, c.expected, actual)
		}
	}
}

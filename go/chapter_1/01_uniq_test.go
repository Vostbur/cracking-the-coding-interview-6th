package chapter_1

import "testing"

func TestIsUnique(t *testing.T) {
	cases := []struct {
		data   string
		result bool
	}{
		{"abcd", true},
		{"abcc", false},
		{" ", true},
		{"", true},
	}
	for _, c := range cases {
		if actual := IsUnique(c.data); actual != c.result {
			t.Errorf("Data: %s. Expected: %v. Actual: %v\n", c.data, c.result, actual)
		}
	}
}

package chapter_1

import "testing"

func TestShift(t *testing.T) {
	cases := []struct {
		input    string
		sub      string
		expected bool
	}{
		{"waterbottle", "terbottlewa", true},
		{"hellomynameis", "nameishellomy", true},
		{"waterbottle", "water", false},
		{"waterbottle", "elttobretaw", false},
		{" ", " ", true},
		{"", "", true},
	}
	for _, c := range cases {
		if actual := Shift(c.input, c.sub); actual != c.expected {
			t.Errorf("Input: (%s, %s). Expected: %v. Actual: %v\n", c.input, c.sub, c.expected, actual)
		}
	}
}

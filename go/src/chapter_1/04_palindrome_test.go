package chapter_1

import "testing"

func TestPalindrome(t *testing.T) {
	cases := []struct {
		input string
		expected bool
	}{
		{"Tact Coa", true},
        {"jhsabckuj ahjsbckj", true},
        {"Able was I ere I saw Elba", true},
        {"So patient a nurse to nurse a patient so", false},
        {"Random Words", false},
        {"Not a Palindrome", false},
        {"no x in nixon", true},
        {"azAZ", true},
	}
	for _, c := range cases {
		if actual := Palindrome(c.input); actual != c.expected {
			t.Errorf("Input: %s. Expected: %v. Actual: %v", c.input, c.expected, actual)
		}
	}
}
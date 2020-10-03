package chapter_1

import (
	"reflect"
	"testing"
)

func TestURLify(t *testing.T) {
	cases := []struct {
		input    string
		length   int
		expected string
	}{
		{"as d  ", 4, "as%20d"},
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
	}
	for _, c := range cases {
		if actual := URLify(c.input, c.length); actual != c.expected {
			t.Errorf("Input: %s %d. Expected: %s. Actual: %s\n", c.input, c.length, c.expected, actual)
		}
	}
}

func TestURLifyInPlace(t *testing.T) {
	cases := []struct {
		input    []rune
		length   int
		expected []rune
	}{
		{[]rune("hello my name is      "), 16, []rune("hello%20my%20name%20is")},
		{[]rune("hello"), 5, []rune("hello")},
		{[]rune(" hello my name is        "), 17, []rune("%20hello%20my%20name%20is")},
	}
	for _, c := range cases {
		URLifyInPlace(c.input, c.length)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("Expected: %s. Actual: %s", string(c.expected), string(c.input))
		}
	}
}

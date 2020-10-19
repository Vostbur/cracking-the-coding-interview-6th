package chapter_3

import (
	"reflect"
	"testing"
)

func TestPlates(t *testing.T) {
	r := GetRowOfPiles(5)
	for i := 1; i < 16; i++ {
		r.push(i)
	}

	expected, actual := []int{}, []int{}
	var value int
	for i := 15; i > 0; i-- {
		expected = append(expected, i)
		value, _ = r.pop()
		actual = append(actual, value)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
}

func TestPlatesPopAt(t *testing.T) {
	r := GetRowOfPiles(5)
	for i := 1; i < 16; i++ {
		r.push(i)
	}

	expected, actual := []int{}, []int{}
	var value int
	for i := 5; i < 16; i++ {
		expected = append(expected, i)
		value = r.popAt(1)
		actual = append(actual, value)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
}

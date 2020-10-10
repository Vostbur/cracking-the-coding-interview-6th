package chapter_2

import (
	"reflect"
	"testing"
)

func TestRemoveDups(t *testing.T) {
	var ll LinkedList
	data := []int{1, 2, 2, 3, 3, 3}
	expect := []int{1, 2, 3}

	ll.AddMultiple(data)
	RemoveDups(&ll)
	if actual := ll.GetSlice(); !reflect.DeepEqual(actual, expect) {
		t.Errorf("Case: %v. Expect: %v. Actual: %v\n", data, expect, actual)
	}
}

func TestRemoveDupsHash(t *testing.T) {
	var ll LinkedList
	data := []int{1, 2, 2, 3, 3, 3}
	expect := []int{1, 2, 3}

	ll.AddMultiple(data)
	RemoveDupsHash(&ll)
	if actual := ll.GetSlice(); !reflect.DeepEqual(actual, expect) {
		t.Errorf("Case: %v. Expect: %v. Actual: %v\n", data, expect, actual)
	}
}

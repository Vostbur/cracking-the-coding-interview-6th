package chapter_2

import (
	"reflect"
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	var ll LinkedList
	expect := []int{1, 2, 3, 5, 6, 7}

	ll.AddMultiple([]int{1, 2, 3, 4, 5, 6, 7})
	DeleteMiddle(&ll, 3)
	if actual := ll.GetSlice(); !reflect.DeepEqual(actual, expect) {
		t.Errorf("Expect: %v. Actual: %v\n", expect, actual)
	}
}

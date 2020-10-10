package chapter_2

import (
	"reflect"
	"testing"
)

func TestDivideByValue(t *testing.T) {
	var ll LinkedList
	data := []int{3, 5, 8, 5, 10, 2, 1}
	expect := []int{1, 2, 3, 5, 8, 5, 10}

	ll.AddMultiple(data)
	result := DivideByValue(ll, 5)
	if actual := result.GetSlice(); !reflect.DeepEqual(actual, expect) {
		t.Errorf("Case: %v. Divide by: %d. Expect: %v. Actual: %v\n",
			data, 5, expect, actual)
	}
}

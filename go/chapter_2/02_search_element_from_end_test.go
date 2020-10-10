// Счет элемента с конца начинаю с 1
package chapter_2

import (
	"testing"
)

func TestGetElementFromEnd(t *testing.T) {
	var ll LinkedList
	data := []int{1, 2, 3, 4, 5, 6}
	expect := 3

	ll.AddMultiple(data)
	if actual := GetElementFromEnd(ll, 4); actual != expect {
		t.Errorf("Case: %v. Element: %d. Expect: %d. Actual: %d\n",
			data, 4, expect, actual)
	}
}

func TestGetElementFromEnd2(t *testing.T) {
	var ll LinkedList
	data := []int{1, 2, 3, 4, 5, 6}
	expect := 3

	ll.AddMultiple(data)
	if actual := GetElementFromEnd2(ll, 4); actual != expect {
		t.Errorf("Case: %v. Element: %d. Expect: %d. Actual: %d\n",
			data, 4, expect, actual)
	}
}

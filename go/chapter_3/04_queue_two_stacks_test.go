package chapter_3

import (
	"reflect"
	"testing"
)

func TestQueueFromStacks(t *testing.T) {
	q := GetQueueFromStacks()

	init := []int{1, 2, 3, 4, 5}
	for _, i := range init {
		q.push(i)
	}

	if peeked, _ := q.peek(); peeked != 1 {
		t.Errorf("Expected: 1. Peeked: %d", peeked)
	}

	actual := []int{}
	for {
		i, err := q.pop()
		if err != nil {
			break
		}
		actual = append(actual, i)
	}

	if !reflect.DeepEqual(init, actual) {
		t.Errorf("Expected: %v. Actual: %v", init, actual)
	}
}

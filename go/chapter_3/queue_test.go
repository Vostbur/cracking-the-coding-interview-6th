package chapter_3

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	init := []int{1, 2, 3, 4, 5}
	q := Queue{}

	for _, i := range init {
		q.Add(i)
	}

	fifo := init[0]
	if r, _ := q.Peek(); r != fifo {
		t.Errorf("Expect: %d. Actual: %d\n", fifo, r)
	}

	actual := []int{}
	for {
		r, err := q.Remove()
		if err != nil {
			break
		}
		actual = append(actual, r)
	}
	if !reflect.DeepEqual(actual, init) {
		t.Errorf("Expect: %v. Actual: %v\n", init, actual)
	}

	if r := q.IsEmpty(); r != true {
		t.Errorf("Expect: %v. Actual: %v\n", true, r)
	}

	if _, err := q.Peek(); err == nil {
		t.Errorf("Queue is not empty.")
	}
}

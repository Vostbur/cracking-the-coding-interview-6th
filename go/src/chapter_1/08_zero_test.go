package chapter_1

import (
	"reflect"
	"testing"
)

func TestZero(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 0, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{1, 0, 3},
				[]int{0, 0, 0},
				[]int{7, 0, 9},
			},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 0},
			},
			[][]int{
				[]int{1, 2, 0},
				[]int{4, 5, 0},
				[]int{0, 0, 0},
			},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{0, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{0, 2, 3},
				[]int{0, 0, 0},
				[]int{0, 8, 9},
			},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
		},
	}
	for _, c := range cases {
		if Zero(c.input); !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("Expected: %v. Actual: %v\n", c.expected, c.input)
		}
	}
}

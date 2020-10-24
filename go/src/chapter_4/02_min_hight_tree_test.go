package chapter_4

import "testing"

func TestSliceToTree(t *testing.T) {
	init := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 18, 22, 222, 2222, 22222}
	tree := initSlice(init)
	tree.print(0, 0)
}

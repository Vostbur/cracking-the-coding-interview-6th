package chapter_4

import "testing"

func TestCountPaths(t *testing.T) {
	tree := GetBinaryTree(50)
	tree.Insert(-20).
		Insert(-10).
		Insert(-10).
		Insert(-20).
		Insert(20).
		Insert(55).
		Insert(8).
		Insert(15).
		Insert(5).
		Insert(-10)
	tree.Print()
	expected := 3
	if actual := countPaths(tree, 10); actual != expected {
		t.Errorf("Expected: %d. Actual: %d", expected, actual)
	}
}

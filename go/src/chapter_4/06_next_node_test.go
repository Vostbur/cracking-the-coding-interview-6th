package chapter_4

import "testing"

func TestNextNode(t *testing.T) {
	tree := GetBinaryTree(50)
	tree.Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)
	expected := -10
	if actual := nextNode(tree.left.right).value; actual != expected {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
}

package chapter_4

import "testing"

func TestCommonAncestor(t *testing.T) {
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
	tree.Print()
	expected := tree
	if actual := commonAncestor(tree.left.left, tree.right.right); actual != expected {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
}

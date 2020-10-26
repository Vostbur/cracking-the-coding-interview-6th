package chapter_4

import "testing"

func TestIsBalancedFalse(t *testing.T) {
	tree := GetBinaryTree(50)
	tree.Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)
	tree.Print()
	
	expected := false
	if actual := isBalanced(tree); actual {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
}

func TestIsBalancedTrue(t *testing.T) {
	tree := GetBinaryTree(50)
	tree.Insert(-30).
		Insert(-20).
		Insert(-40).
		Insert(50).
		Insert(70).
		Insert(60).
		Insert(90)
	tree.Print()
	
	expected := true
	if actual := isBalanced(tree); !actual {
		t.Errorf("Expected: %v. Actaul: %v", expected, actual)
	}
}

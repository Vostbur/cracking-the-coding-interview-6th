package chapter_4

import "testing"

func TestCheckBST(t *testing.T) {
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
	expected := true
	if actual := checkBST(tree); !actual {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
	
	tree.left.right.value = -25
	expected = false
	if actual := checkBST(tree); actual {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}

	if actual := checkBST2(tree); actual {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}

	tree.left.right.value = -15
	expected = true
	if actual := checkBST2(tree); !actual {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
}
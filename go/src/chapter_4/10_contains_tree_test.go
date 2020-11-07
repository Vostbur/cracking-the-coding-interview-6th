package chapter_4

import "testing"

func TestContainsTree(t *testing.T) {
	tree1 := GetBinaryTree(50)
	tree1.Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)

	tree2 := GetBinaryTree(60)
	tree2.Insert(55).Insert(85)

	if actual := containsTree(tree1, tree2); !actual {
		t.Errorf("Expected 'true', actual %v", actual)
	}

	tree2.Insert(95)
	if actual := containsTree(tree1, tree2); actual {
		t.Errorf("Expected 'false', actual %v", actual)
	}
}

package chapter_4

import "testing"

func TestRandomNode(t *testing.T) {
	n := getTreeNode(20)
	init := []int{10, 30, 5, 15, 3, 7, 17}
	for _, i := range init {
		n.insertInOrder(i)
	}

	actual := n.getRandomNode()
	if expected := n.find(actual.data); actual != expected {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
}

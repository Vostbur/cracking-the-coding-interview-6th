package chapter_4

import (
	"fmt"
)

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func GetBinaryTree(value int) *BinaryTree {
	return &BinaryTree{value: value}
}

func (t *BinaryTree) Insert(value int) *BinaryTree {
	if t.value >= value {
		if t.left == nil {
			t.left = &BinaryTree{value: value}
			return t
		}
		t.left.Insert(value)
	}

	if t.value < value {
		if t.right == nil {
			t.right = &BinaryTree{value: value}
			return t
		}
		t.right.Insert(value)
	}
	return t
}

func (t *BinaryTree) Remove(value int) *BinaryTree {
	if t == nil {
		return nil
	}

	if value < t.value {
		t.left = t.left.Remove(value)
		return t
	}
	if value > t.value {
		t.right = t.right.Remove(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestOnRight := t.right
	for {
		if smallestOnRight != nil && smallestOnRight.left != nil {
			smallestOnRight = smallestOnRight.left
		} else {
			break
		}
	}

	t.value = smallestOnRight.value
	t.right = t.right.Remove(t.value)
	return t
}

func (t *BinaryTree) GetNode(value int) (BinaryTree, bool) {
	if t == nil {
		return BinaryTree{}, false
	}

	switch {
	case value == t.value:
		return *t, true
	case value < t.value:
		return t.left.GetNode(value)
	default:
		return t.right.GetNode(value)
	}
}

func (t BinaryTree) Max() int {
	if t.right == nil {
		return t.value
	}
	return t.right.Max()
}

func (t BinaryTree) Min() int {
	if t.left == nil {
		return t.value
	}
	return t.left.Min()
}

func (t BinaryTree) Print() {
	t.printTree(0, 0)
}

func (t *BinaryTree) printTree(level int, ch rune) {
	if t != nil {
		t.left.printTree(level+1, '┌')
		for i := 0; i < level*4; i++ {
			fmt.Print("  ")
		}
		fmt.Printf("%c---", ch)
		fmt.Printf("-> %d\n", t.value)
		t.right.printTree(level+1, '└')
	}
}

package chapter_4

import (
	"fmt"
	"testing"
)

func TestAllSequences(t *testing.T) {
	tree := GetBinaryTree(3)
	tree.Insert(1).
		Insert(2).
		Insert(4).
		Insert(5)
	tree.Print()

	allSeq := allSequences(tree)
	for _, l := range allSeq {
		fmt.Printf("%v\n", l.GetSlice())
	}
}

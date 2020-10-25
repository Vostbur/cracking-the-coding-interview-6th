package chapter_4

import (
	"fmt"
	"testing"
)

func TestListsFromTree(t *testing.T) {
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
	lists := tree.createListsDFS()
	for i, l := range lists {
		fmt.Printf("Level: %d --> %v\n", i, l.GetSlice())
	}
}

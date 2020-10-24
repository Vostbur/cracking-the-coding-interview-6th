package chapter_4

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
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
		Insert(-10).
		Remove(85).
		Insert(85)
	tree.Print()
	tree.Remove(60)
	tree.Print()
	tree.Remove(-20)
	tree.Print()
	fmt.Println(tree.Max())
	fmt.Println(tree.Min())
	if n, check := tree.GetNode(-15); check {
		fmt.Printf("%v", n)
	}
}

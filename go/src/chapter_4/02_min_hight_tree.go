// Напишите алгоритм создания бинарного дерева поиска с минимальной
// высотой для отсортированного (по возрастанию) массива с уникальными
// целочисленными элементами.
package chapter_4

import "fmt"

type MinBinaryTree struct {
	value int
	left  *MinBinaryTree
	right *MinBinaryTree
}

func (t *MinBinaryTree) print(level int, ch rune) {
	if t != nil {
		t.left.print(level+1, '┌')
		for i := 0; i < level*4; i++ {
			fmt.Print("  ")
		}
		fmt.Printf("%c---", ch)
		fmt.Printf("-> %d\n", t.value)
		t.right.print(level+1, '└')
	}
}

func initSlice(s []int) *MinBinaryTree {
	return sliceToTree(s, 0, len(s)-1)
}

func sliceToTree(s []int, start int, end int) *MinBinaryTree {
	if start > end {
		return nil
	}
	mid := int((start + end) / 2)
	node := &MinBinaryTree{value: s[mid]}
	node.left = sliceToTree(s, start, mid-1)
	node.right = sliceToTree(s, mid+1, end)
	return node
}

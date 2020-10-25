// Реализуйте функцию, проверяющую сбалансированность бинарного дерева.
// Будем считать дерево сбалансированным, если разница высот двух поддеревьев любого узла не превышает 1.
package chapter_4

import "math"

func isBalanced(tree *BinaryTree) bool {
	return checkHeight(tree) != math.MinInt64
}

func checkHeight(tree *BinaryTree) int {
	if tree == nil {
		return -1
	}

	leftHeight := checkHeight(tree.left)
	if leftHeight == math.MinInt64 {
		return math.MinInt64
	}

	rightHeight := checkHeight(tree.right)
	if rightHeight == math.MinInt64 {
		return math.MinInt64
	}

	heightDiff := leftHeight - rightHeight
	if math.Abs(float64(heightDiff)) > 1 {
		return math.MinInt64
	} else {
		return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
	}
}

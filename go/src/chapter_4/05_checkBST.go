// Реализуйте функцию для проверки того, является ли бинарное дерево бинарным деревом поиска.
package chapter_4

func checkBST(n *BinaryTree) bool {
	var last_printed *int
	return checkBst(n, last_printed)
}

// Do not work with duplicates
func checkBst(n *BinaryTree, last_printed *int) bool {
	if n == nil {
		return true
	}
	if !checkBst(n.left, last_printed) {
		return false
	}
	if last_printed != nil && n.value <= *last_printed {
		return false
	}
	last_printed = &n.value
	return checkBst(n.right, last_printed)
	// the same:
	// if !checkBst(n.right, last_printed) {
	// 	return false
	// }
	// return true
}


// second version (work with duplicates)
func checkBST2(n *BinaryTree) bool {
	return checkBSTMinMax(n, nil, nil)
}

func checkBSTMinMax(n *BinaryTree, min, max *int) bool {
	if n == nil {
		return true
	}
	if (min != nil && n.value <= *min) || (max != nil && n.value > *max) {
		return false
	}
	if !checkBSTMinMax(n.left, min, &n.value) || !checkBSTMinMax(n.right, &n.value, max) {
		return false
	}
	return true
}

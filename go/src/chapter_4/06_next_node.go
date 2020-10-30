// Напишите алгоритм поиска «следующего» узла для заданного узла в бинарном
// дереве поиска. Считайте, что у каждого узла есть ссылка на его родителя.
package chapter_4

func nextNode(n *BinaryTree) *BinaryTree {
	if n == nil {
		return nil
	}
	if n.right != nil {
		return leftChild(n.right)
	} else {
		q := n
		x := q.parent
		for x != nil && x.left != q {
			q = x
			x = x.parent
		}
		return x
	}
}

func leftChild(n *BinaryTree) *BinaryTree {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

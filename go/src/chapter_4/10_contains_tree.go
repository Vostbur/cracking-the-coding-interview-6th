// T1 и T2 — два очень больших бинарных дерева, причем T1 значительно больше T2.
// Создайте алгоритм, проверяющий, является ли T2 поддеревом T1.
// Дерево T2 считается поддеревом T1, если существует такой узел n в T1,
// что поддерево, «растущее» из n, идентично дереву T2.
// (Иначе говоря, если вырезать дерево в узле n, оно будет идентично T2.)
package chapter_4

func containsTree(t1, t2 *BinaryTree) bool {
	// empty tree is a subtree for any tree
	if t2 == nil {
		return true
	}
	return subTree(t1, t2)
}

func subTree(r1, r2 *BinaryTree) bool {
	if r1 == nil {
		return false
	}
	if r1.value == r2.value && matchTree(r1, r2) {
		return true
	}
	return subTree(r1.left, r2) || subTree(r1.right, r2)
}

func matchTree(r1, r2 *BinaryTree) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil || r1.value != r2.value {
		return false
	}
	return matchTree(r1.left, r2.left) && matchTree(r1.right, r2.right)
}

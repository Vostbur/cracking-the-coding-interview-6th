// Создайте алгоритм и напишите код поиска первого общего предка двух узлов
// бинарного дерева. Постарайтесь избежать хранения дополнительных узлов
// в структуре данных. Примечание: бинарное дерево не обязательно является
// бинарным деревом поиска.
package chapter_4

func commonAncestor(p, q *BinaryTree) *BinaryTree {
	delta := depth(p) - depth(q)
	first, second := p, q
	if delta < 0 {
		first = q
		second = p
		delta = -delta
	}
	second = goUpBy(second, delta)
	for first != second && first != nil && second != nil {
		first = first.parent
		second = second.parent
	}
	if first == nil || second == nil {
		return nil
	}
	return first
}

func depth(n *BinaryTree) int {
	d := 0
	for n != nil {
		n = n.parent
		d++
	}
	return d
}

func goUpBy(n *BinaryTree, delta int) *BinaryTree {
	for delta > 0 && n != nil {
		n = n.parent
		delta--
	}
	return n
}

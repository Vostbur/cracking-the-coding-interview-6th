// Бинарное дерево поиска было создано обходом массива слева направо и вставкой каждого элемента.
// Для заданного бинарного дерева поиска с разными элементами выведите все возможные массивы,
// которые могли привести к созданию этого дерева.
// Пример:
// Ввод: 2
//      1 3
// Вывод: {2, 1, 3}, {2, 3, 1}
package chapter_4

func allSequences(node *BinaryTree) []*LinkedList {
	
	// получаем связанный список узлов левого поддерева 
	leftSeq := GetLinkedList()
	inOrderTraversal(node.left, leftSeq)
	// получаем связанный список узлов правого поддерева 
	rightSeq := GetLinkedList()
	inOrderTraversal(node.right, rightSeq)

	// переплетаем два списка с сохранением порядка следования элементов
	weaved := []*LinkedList{}
	weaveLists(leftSeq, rightSeq, &weaved, GetLinkedList())
	
	// в результат каждого переплетения добавляем корневой элемент
	for _, w := range weaved {
		w.AddToHead(node.value)
	}
	return weaved
}

func inOrderTraversal(n *BinaryTree, l *LinkedList) {
	if n != nil {
		inOrderTraversal(n.left, l)
		l.Add(n.value)
		inOrderTraversal(n.right, l)
	}
}

// Алгоритм работы для переплетения {1, 2, 3} с {4, 5, 6}
// weave(first, second, prefix):
// 	weave({1, 2}, {3, 4}, {})
// 	 weave({2}, {3, 4}, {1})
//    weave({}, {3, 4}, {1, 2})
//     {1, 2, 3, 4}
//    weave({2}, {4}, {1, 3})
//     weave({}, {4}, {1, 3, 2})
//  	{1, 3, 2, 4}
//     weave({2}, {}, {1, 3, 4})
//      {1, 3, 4, 2}
//  weave({1, 2}, {4}, {3})
//   weave({2}, {4}, {3, 1})
//    weave({}, {4}, {3, 1, 2})
//     {3, 1, 2, 4}
//    weave({2}, {}, {3, 1, 4})
//     {3, 1, 4, 2}
//   weave({1, 2}, {}, {3, 4})
//    {3, 4, 1, 2}
func weaveLists(first, second *LinkedList, results *[]*LinkedList, prefix *LinkedList) {
	if first.Len() == 0 || second.Len() == 0 {
		result := prefix.Clone()
		result.AddAll(first)
		result.AddAll(second)
		*results = append(*results, result)
		return
	}

	headFirst, _ := first.GetNodeByIndex(0)
	_ = first.DelByIndex(0)
	valHeadFirst := headFirst.value
	prefix.Add(valHeadFirst)
	weaveLists(first, second, results, prefix)
	_ = prefix.DelByIndex(prefix.Len() - 1)
	first.AddToHead(valHeadFirst)

	headSecond, _ := second.GetNodeByIndex(0)
	_ = second.DelByIndex(0)
	valHeadSecond := headSecond.value
	prefix.Add(valHeadSecond)
	weaveLists(first, second, results, prefix)
	_ = prefix.DelByIndex(prefix.Len() - 1)
	second.AddToHead(valHeadSecond)
}

func (l *LinkedList) Clone() *LinkedList {
	return GetLinkedListFromSlice(l.GetSlice())
}

func (l *LinkedList) AddAll(list *LinkedList) {
	l.AddMultiple(list.GetSlice())
}

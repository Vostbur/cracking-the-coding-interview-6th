// Для бинарного дерева разработайте алгоритм, который создает связный
// список всех узлов, находящихся на каждой глубине (для дерева с глубиной
// D должно получиться D связных списков).
package chapter_4

// depth-first search case
func (t *BinaryTree) createListsDFS() []*LinkedList {
	lists := []*LinkedList{}
	t.createLevelLinkedListDFS(&lists, 0)
	return lists
}

// Модифицированный вариант поиска в глубину
func (t *BinaryTree) createLevelLinkedListDFS(lists *[]*LinkedList, level int) {
	if t == nil {
		return
	}
	if len(*lists) == level {
		// Если это новый уровень, создаем новый связанный список для хранения 
		// узлов уровня и добавляем в него первый узел
		list := GetLinkedList()
		list.Add(t.value)
		*lists = append(*lists, list)
	} else {
		// Список этого уровня уже существует, добавляем в массив узлов новый узел уровня
		(*lists)[level].Add(t.value)
	}
	t.left.createLevelLinkedListDFS(lists, level+1)
	t.right.createLevelLinkedListDFS(lists, level+1)
}

// Для кольцевого связного списка реализуйте алгоритм, возвращающий начальный узел петли.
// Определение:
// Кольцевой связный список — это связный список, в котором указатель следующего
// узла ссылается на более ранний узел, образуя петлю.
// Пример:
// Ввод: A->B->C->D->E->C (предыдущий узел C)
// Вывод: C
package chapter_2

func (ll *LinkedList) LoopDetect() *node {
	fast, slow := ll.head, ll.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.next == nil {
		return nil
	}

	slow = ll.head
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}

	return fast
}

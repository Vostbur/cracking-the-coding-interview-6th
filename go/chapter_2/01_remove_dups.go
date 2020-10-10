// Напишите код для удаления дубликатов из несортированного связного списка.
// Дополнительно:
// Как вы будете решать задачу, если использовать временный буфер запрещено?
package chapter_2

func RemoveDups(ll *LinkedList) {
	if ll.count < 2 {
		return
	}
	n := ll.head
	for n != nil {
		next_n := n
		for next_n.next != nil {
			if next_n.next.value == n.value {
				next_n.next = next_n.next.next
				ll.count--
			} else {
				next_n = next_n.next
			}
		}
		n = n.next
	}
}

func RemoveDupsHash(ll *LinkedList) {
	if ll.count < 2 {
		return
	}
	uniq := make(map[int]struct{})
	n := ll.head
	for n != nil {
		if _, ok := uniq[n.value]; ok {
			ll.DeleteNode(n)
		} else {
			uniq[n.value] = struct{}{}
		}
		n = n.next
	}
}
// Реализуйте алгоритм для нахождения в односвязном списке k-го элемента с конца
package chapter_2

func GetElementFromEnd(ll LinkedList, x int) int {
	e, _ := ll.GetNodeByIndex(ll.count - x)
	return e.value
}

// Если длина односвязного списка неизвестна
func GetElementFromEnd2(ll LinkedList, x int) int {
	first, second := ll.head, ll.head

	for i := 0; i < x; i++ {
		second = second.next
		if second == nil {
			return -1
		}
	}

	for second != nil {
		first = first.next
		second = second.next
	}

	return first.value
}

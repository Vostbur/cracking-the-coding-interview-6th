// Проверьте, пересекаются ли два заданных (одно-)связных списка.
// Верните узел пересечения. Учтите, что пересечение определяется ссылкой, а не значением.
// Иначе говоря, если k-й узел первого связного списка точно совпадает
// (по ссылке) с j-м узлом второго связного списка, то списки считаются пересекающимися
package chapter_2

func Intersection(l1, l2 *LinkedList) (bool, *node) {
	if l1.tail != l2.tail {
		return false, nil
	}

	diff := l1.Len() - l2.Len()
	n1, n2 := l1.head, l2.head
	if diff > 0 {
		for i := 0; i < diff; i++ {
			n1 = n2.next
		}
	}
	if diff < 0 {
		diff = -diff
		for i := 0; i < diff; i++ {
			n2 = n2.next
		}
	}

	for n1 != nil {
		if n1 == n2 {
			return true, n1
		}
		n1 = n1.next
		n2 = n2.next
	}
	return false, nil
}

// Два числа хранятся в виде связных списков, в которых каждый узел представляет один разряд.
// Все цифры хранятся в обратном порядке, при этом
// младший разряд (единицы) хранится в начале списка. Напишите функцию,
// которая суммирует два числа и возвращает результат в виде связного списка.
// Пример:
// Ввод: (7->1->6) + (5->9->2), то есть 617 + 295.
// Вывод: 2->1->9, то есть 912.
// Дополнительно
// Решите задачу, предполагая, что цифры записаны в прямом порядке.
// Ввод: (6->1->7) + (2->9->5), то есть 617 + 295.
// Вывод: 9->1->2, то есть 912.
package chapter_2

import "strconv"

// Сумма элементов записанных в обратном порядке
func Accumulator(l1, l2 LinkedList) *LinkedList {
	a, b := l1.head, l2.head
	shift, sum := 0, 0
	result := GetLinkedList()

	for a != nil || b != nil {
		sum = shift
		if a != nil {
			sum += a.value
			a = a.next
		}
		if b != nil {
			sum += b.value
			b = b.next
		}
		shift = sum / 10
		result.Add(sum % 10)
	}

	if shift != 0 {
		result.Add(shift)
	}

	return result
}

// Сумма элементов записанных в прямом порядке
func AccumulatorInRow(l1, l2 LinkedList) *LinkedList {
	diff := l1.Len() - l2.Len()
	if diff < 0 {
		diff = -diff
		l1.AddToHeadMultiple(make([]int, diff))
	} else {
		l2.AddToHeadMultiple(make([]int, diff))
	}

	result := GetLinkedList()
	sum := 0
	a, b := l1.head, l2.head
	for a != nil {
		sum = sum*10 + a.value + b.value
		a = a.next
		b = b.next
	}

	for _, i := range []byte(strconv.Itoa(sum)) {
		tmp, _ := strconv.Atoi(string(i))
		result.Add(tmp)
	}

	return result
}

// Реализуйте функцию, проверяющую, является ли связный список палиндромом.
package chapter_2

// Для двусвязного списка
func (ll *LinkedList) IsPalindrome() bool {
	if ll.head == nil || ll.head == ll.tail {
		return true
	}
	h, t := ll.head, ll.tail
	for {
		if h.value != t.value {
			return false
		}
		h = h.next
		t = t.prev
		if h == t || h.prev == t {
			break
		}
	}
	return true
}

// Для односвязного списка с использованием стека
func (ll *LinkedList) IsPalindromeStack() bool {
	left, right := ll.head, ll.head
	stack := []int{}

	for right != nil && right.next != nil {
		stack = append(stack, left.value)
		left = left.next
		right = right.next.next
	} 

	if right != nil {
		left = left.next
	}

	for left != nil {
		if i := pop(&stack); i != left.value {
			return false
		}
		left = left.next
	}

	return true
}

func pop(stack *[]int) int {
	last := len(*stack) - 1
	item := (*stack)[last]
	*stack = (*stack)[:last]
	return item
}

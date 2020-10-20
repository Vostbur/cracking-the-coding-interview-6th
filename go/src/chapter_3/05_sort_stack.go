// Напишите программу сортировки стека, в результате которой наименьший
// элемент оказывается на вершине стека. Вы можете использовать дополнительный временный стек,
// но элементы не должны копироваться в другие
// структуры данных (например, в массив). Стек должен поддерживать следующие операции: push, pop, peek, isEmpty.
package chapter_3

type SortedStack struct {
	sortStack *Stack
	tempStack *Stack
}

func GetSortedStack() *SortedStack {
	return &SortedStack{GetStack(), GetStack()}
}

func (s *SortedStack) push(value int) {
	// pop values until empty or fine the larger one
	for {
		peeked, err := s.sortStack.peek()
		if err != nil {
			break
		}
		if peeked < value {
			poped, _ := s.sortStack.pop()
			s.tempStack.push(poped)
		} else {
			break
		}
	}
	s.sortStack.push(value)
	// put back smaller values
	for {
		poped, err := s.tempStack.pop()
		if err != nil {
			break
		}
		s.sortStack.push(poped)
	}
}

func (s *SortedStack) pop() (int, error) {
	return s.sortStack.pop()
}

func (s *SortedStack) peek() (int, error) {
	return s.sortStack.peek()
}

func (s *SortedStack) isEmpty() bool {
	return s.sortStack.isEmpty()
}

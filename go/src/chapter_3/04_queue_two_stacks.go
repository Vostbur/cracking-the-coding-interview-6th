
// Напишите класс MyQueue, который реализует очередь с использованием двух стеков.
package chapter_3

type QueueFromStacks struct {
	main *Stack
	temp *Stack
}

func GetQueueFromStacks() *QueueFromStacks {
	return &QueueFromStacks{GetStack(), GetStack()}
}

func (q *QueueFromStacks) push(value int) {
	q.main.push(value)
}

func (q *QueueFromStacks) pop() (int, error) {
	swapValues(q.main, q.temp)
	ret, err := q.temp.pop()
	if err != nil {
		return -1, err
	}
	swapValues(q.temp, q.main)
	return ret, nil
}

func (q *QueueFromStacks) peek() (int, error) {
	swapValues(q.main, q.temp)
	ret, err := q.temp.peek()
	if err != nil {
		return -1, err
	}
	swapValues(q.temp, q.main)
	return ret, nil
}

func swapValues(s1, s2 *Stack) {
	for {
		i, err := s1.pop()
		if err != nil {
			break
		}
		s2.push(i)
	}
}

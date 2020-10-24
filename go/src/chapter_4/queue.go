package chapter_4

import "errors"

type queueNode struct {
	value      int
	prev, next *queueNode
}

type Queue struct {
	head, tail *queueNode
}

func GetQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Add(x int) {
	newNode := &queueNode{value: x}
	if q.tail != nil {
		newNode.prev = q.tail
		q.tail.next = newNode
	}
	q.tail = newNode
	if q.head == nil {
		q.head = newNode
	}
}

func (q *Queue) Remove() (int, error) {
	if q.head == nil {
		return -1, errors.New("Cannot Remove. Queue is empty.")
	}
	r := q.head.value
	q.head = q.head.next
	if q.head != nil {
		q.head.prev = nil
	}
	return r, nil
}

func (q Queue) Peek() (int, error) {
	if q.head == nil {
		return -1, errors.New("Cannot Peek. Queue is empty.")
	}
	return q.head.value, nil
}

func (q Queue) IsEmpty() bool {
	return q.head == nil
}

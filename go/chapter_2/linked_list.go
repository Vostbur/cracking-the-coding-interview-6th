// LinkedList.go
package main

import (
	"errors"
	"fmt"
)

type node struct {
	value int
	next  *node
	prev  *node
}

type LinkedList struct {
	head  *node
	tail  *node
	count int
}

func (ll *LinkedList) Add(x int) {
	newNode := &node{value: x}
	ll.AddNode(newNode)
}

func (ll *LinkedList) AddNode(newNode *node) {
	if ll.count == 0 {
		ll.head, ll.tail = newNode, newNode
	} else {
		ll.tail.next = newNode
		newNode.prev = ll.tail
		ll.tail = newNode
	}
	ll.count++
}

func (ll *LinkedList) AddMultiple(args []int) {
	for _, arg := range args {
		ll.Add(arg)
	}
}

func (ll *LinkedList) AddToHead(x int) {
	if ll.count == 0 {
		ll.Add(x)
	} else {
		newNode := &node{value: x}
		ll.AddNodeToHead(newNode)
		ll.count++
	}
}

func (ll *LinkedList) AddNodeToHead(newNode *node) {
	n := ll.head
	ll.head = newNode
	ll.head.next = n
}

func (ll *LinkedList) AddToHeadMultiple(args []int) {
	for i := len(args) - 1; i >= 0; i-- {
		ll.AddToHead(args[i])
	}
}

func (ll LinkedList) GetSlice() []int {
	s := make([]int, 0, ll.count)
	n := ll.head
	for i := 0; i < ll.count; i++ {
		s = append(s, n.value)
		n = n.next
	}
	return s
}

func (ll LinkedList) GetReverseSlice() []int {
	s := make([]int, 0, ll.count)
	n := ll.tail
	for i := 0; i < ll.count; i++ {
		s = append(s, n.value)
		n = n.prev
	}
	return s
}

func (ll LinkedList) GetNode(x int) (*node, error) {
	n := ll.head
	for i := 0; i < ll.count; i++ {
		if n.value == x {
			return n, nil
		}
		n = n.next
	}
	return nil, errors.New(fmt.Sprintf("linked list has no element '%d'\n", x))
}

func (ll LinkedList) GetNodeByIndex(x int) (*node, error) {
	n := ll.head
	for i := 0; i < ll.count; i++ {
		if i == x {
			return n, nil
		}
		n = n.next
	}
	return nil, errors.New(fmt.Sprintf("linked list has no index '%d'\n", x))
}

func (ll *LinkedList) Del(x int) error {
	n, err := ll.GetNode(x)
	if err != nil {
		return err
	}
	ll.DeleteNode(n)
	return nil
}

func (ll *LinkedList) DelByIndex(x int) error {
	n, err := ll.GetNodeByIndex(x)
	if err != nil {
		return err
	}
	ll.DeleteNode(n)
	return nil

}

func (ll *LinkedList) DeleteNode(n *node) {
	ll.count--
	if ll.count == 0 {
		ll.head, ll.tail = nil, nil
		return
	}
	if n.prev == nil {
		ll.head = ll.head.next
		ll.head.prev = nil
		return
	}
	if n.next == nil {
		ll.tail = ll.tail.prev
		ll.tail.next = nil
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (ll LinkedList) Len() int {
	return ll.count
}

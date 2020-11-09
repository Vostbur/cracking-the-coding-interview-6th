// Вы пишете с нуля класс бинарного дерева поиска, который помимо методов
// вставки, поиска и удаления содержит метод getRandomNode() для получения
// случайного узла дерева. Вероятность выбора всех узлов должна быть одинаковой.
// Разработайте и реализуйте алгоритм getRandomNode; объясните, как вы
// реализуете остальные методы.
package chapter_4

import (
	"time"
	"math/rand"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
	size  int
}

func getTreeNode(data int) *TreeNode {
	return &TreeNode{data: data, size: 1}
}

func (n *TreeNode) getRandomNode() *TreeNode {
	var leftSize int
	if n.left == nil {
		leftSize = 0
	} else {
		leftSize = n.left.size 
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(n.size)
	if index < leftSize {
		return n.left.getRandomNode()
	} else if index == leftSize {
		return n
	} else {
		return n.right.getRandomNode()
	}
}

func (n *TreeNode) insertInOrder(d int) {
	if d <= n.data {
		if n.left == nil {
			n.left = getTreeNode(d)
		} else {
			n.left.insertInOrder(d)
		}
	} else {
		if n.right == nil {
			n.right = getTreeNode(d)
		} else {
			n.right.insertInOrder(d)
		}
	}
	n.size++
}

func (n *TreeNode) find(d int) *TreeNode {
	if d == n.data {
		return n
	} else if d <= n.data {
		if n.left != nil {
			return n.left.find(d)
		} else {
			return nil
		}
	} else if d > n.data {
		if n.right != nil {
			return n.right.find(d)
		} else {
			return nil
		}
	} 
	return nil
}

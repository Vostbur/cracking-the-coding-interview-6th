# Вы пишете с нуля класс бинарного дерева поиска, который помимо методов
# вставки, поиска и удаления содержит метод getRandomNode() для получения
# случайного узла дерева. Вероятность выбора всех узлов должна быть одинаковой.
# Разработайте и реализуйте алгоритм getRandomNode; объясните, как вы
# реализуете остальные методы.
import random
import unittest


class TreeNode:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
        self.size = 1

    @property
    def get_random_node(self):
        left_size = 0 if not self.left else self.left.size
        index = random.randint(0, self.size-1)
        if index < left_size:
            return self.left.get_random_node
        elif index == left_size:
            return self
        else:
            return self.right.get_random_node

    def insert_in_order(self, d):
        if d <= self.data:
            if not self.left:
                self.left = TreeNode(d)
            else:
                self.left.insert_in_order(d)
        else:
            if not self.right:
                self.right = TreeNode(d)
            else:
                self.right.insert_in_order(d)
        self.size += 1

    def find(self, d):
        if d == self.data:
            return self
        elif d <= self.data:
            return self.left.find(d) if self.left else None
        elif d > self.data:
            return self.right.find(d) if self.right else None
        return None


class Test(unittest.TestCase):
    init = [10, 30, 5, 15, 3, 7, 17]

    def test_random_node(self):
        n = TreeNode(20)
        for i in self.init:
            n.insert_in_order(i)
        actual = n.get_random_node
        expected = n.find(actual.data)
        self.assertEqual(actual, expected)


if __name__ == '__main__':
    unittest.main()

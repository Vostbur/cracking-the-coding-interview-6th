# Создайте алгоритм и напишите код поиска первого общего предка двух узлов
# бинарного дерева. Постарайтесь избежать хранения дополнительных узлов
# в структуре данных. Примечание: бинарное дерево не обязательно является
# бинарным деревом поиска.
import unittest
from chapter_4.binary_tree import BinaryTree


def common_ancestor(p: BinaryTree, q: BinaryTree) -> BinaryTree:
    delta = depth(p) - depth(q)
    first = q if delta > 0 else p
    second = p if delta > 0 else q
    second = go_up(second, abs(delta))
    while first != second and first and second:
        first = first.parent
        second = second.parent
    return None if not first or not second else first


def depth(n: BinaryTree) -> int:
    d = 0
    while n:
        n = n.parent
        d += 1
    return d


def go_up(n: BinaryTree, delta: int) -> BinaryTree:
    while delta > 0 and n:
        n = n.parent
        delta -= 1
    return n


class Test(unittest.TestCase):
    init = [-20, -50, -15, -60, 60, 55, 85, 15, 5, -10]
    tree = BinaryTree(50)

    def test_common_ancestor(self):
        for i in self.init:
            self.tree.insert(i)
        self.assertEqual(
            self.tree, common_ancestor(
                self.tree.left.left, self.tree.right.right)
        )


if __name__ == "__main__":
    unittest.main()

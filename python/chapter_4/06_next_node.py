# Напишите алгоритм поиска «следующего» узла для заданного узла в бинарном
# дереве поиска. Считайте, что у каждого узла есть ссылка на его родителя.
import unittest
from chapter_4.binary_tree import BinaryTree


def next_node(node: BinaryTree) -> BinaryTree:
    if not node:
        return None
    if node.right:
        return left_child(node.right)
    else:
        q = node
        x = q.parent
        while x and x.left != q:
            q = x
            x = x.parent
        return x


def left_child(node: BinaryTree) -> BinaryTree:
    if not node:
        return None
    while node.left:
        node = node.left
    return node


class Test(unittest.TestCase):
    init = [-20, -50, -15, -60, 60, 55, 85, 15, 5, -10]
    tree = BinaryTree(50)

    def test_nextNode(self):
        for i in self.init:
            self.tree.insert(i)
        print(self.tree)
        self.assertEqual(-10, next_node(self.tree.left.right).value)


if __name__ == '__main__':
    unittest.main()

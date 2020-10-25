# Реализуйте функцию, проверяющую сбалансированность бинарного дерева.
# Будем считать дерево сбалансированным, если разница высот двух поддеревьев любого узла не превышает 1.
import sys
import unittest
from chapter_4.binary_tree import BinaryTree


MININT = -sys.maxsize - 1


def is_balanced(tree: BinaryTree) -> bool:
    return check_height(tree) != MININT


def check_height(tree: BinaryTree) -> int:
    if not tree:
        return -1

    left_height = check_height(tree.left)
    if left_height == MININT:
        return MININT

    right_height = check_height(tree.right)
    if right_height == MININT:
        return MININT

    height_diff = left_height - right_height
    if abs(height_diff) > 1:
        return MININT
    else:
        return max(left_height, right_height) + 1


class Test(unittest.TestCase):
    def test_is_balanced_false(self):
        init = [-20, -50, -15, -60, 50, 60, 55, 85, 15, 5, -10]
        tree = BinaryTree(50)
        for i in init:
            tree.insert(i)
        print(tree)
        self.assertFalse(is_balanced(tree))

    def test_is_balanced_true(self):
        init = [-30, -20, -40, 50, 70, 60, 90]
        tree = BinaryTree(50)
        for i in init:
            tree.insert(i)
        print(tree)
        self.assertTrue(is_balanced(tree))


if __name__ == '__main__':
    unittest.main()

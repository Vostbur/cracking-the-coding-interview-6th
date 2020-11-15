# Реализуйте функцию для проверки того, является ли бинарное дерево бинарным
# деревом поиска.
import unittest
from chapter_4.binary_tree import BinaryTree


last_printed = None


# do not work with duplicates in the tree
def checkBST(node: BinaryTree) -> bool:
    global last_printed
    if not node:
        return True
    if not checkBST(node.left):
        return False
    if not(last_printed is None) and node.value <= last_printed:
        return False
    last_printed = node.value
    return checkBST(node.right)
    # the same:
    # if not checkBST(node.right):
    #     return False
    # return True


# second version (work with duplicates)
def checkBST_min_max(node: BinaryTree) -> bool:
    return checkBST2(node, None, None)


def checkBST2(node: BinaryTree, min_value: int, max_value: int) -> bool:
    if not node:
        return True
    if (not(min_value is None) and node.value <= min_value) or \
            (not (max_value is None) and node.value > max_value):
        return False
    if not checkBST2(node.left, min_value, node.value) or \
            not checkBST2(node.right, node.value, max_value):
        return False
    return True


class Test(unittest.TestCase):
    init = [-20, -50, -15, -60, 60, 55, 85, 15, 5, -10]
    tree = BinaryTree(50)

    def test_01_checkBST_true(self):
        for i in self.init:
            self.tree.insert(i)
        self.assertTrue(checkBST(self.tree))

    def test_02_checkBST_false(self):
        self.tree.insert(50)
        self.assertFalse(checkBST(self.tree))

    def test_03_checkBST_min_max_true(self):
        self.assertTrue(checkBST_min_max(self.tree))

    def test_04_checkBST_min_max_false(self):
        self.tree.left.right.value = -25
        self.assertFalse(checkBST_min_max(self.tree))


if __name__ == '__main__':
    unittest.main()

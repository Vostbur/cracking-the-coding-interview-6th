# T1 и T2 — два очень больших бинарных дерева, причем T1 значительно больше T2.
# Создайте алгоритм, проверяющий, является ли T2 поддеревом T1.
# Дерево T2 считается поддеревом T1, если существует такой узел n в T1,
# что поддерево, «растущее» из n, идентично дереву T2.
# (Иначе говоря, если вырезать дерево в узле n, оно будет идентично T2.)
import unittest
from chapter_4.binary_tree import BinaryTree


def contains_tree(t1: BinaryTree, t2: BinaryTree) -> bool:
    # empty tree is a subtree for any tree
    if not t1:
        return True
    return sub_tree(t1, t2)


def sub_tree(r1: BinaryTree, r2: BinaryTree) -> bool:
    if not r1:
        return False
    if r1.value == r2.value and match_tree(r1, r2):
        return True
    return sub_tree(r1.left, r2) or sub_tree(r1.right, r2)


def match_tree(r1: BinaryTree, r2: BinaryTree) -> bool:
    if not r1 and not r2:
        return True
    if not r1 or not r2 or r1.value != r2.value:
        return False
    return match_tree(r1.left, r2.left) and match_tree(r1.right, r2.right)


class Test(unittest.TestCase):
    init1, init2 = [-20, -50, -15, -60, 60, 55, 85, 15, 5, -10], [55, 85]
    tree1, tree2 = BinaryTree(50), BinaryTree(60)

    def test_contains_tree_01(self):
        for i in self.init1:
            self.tree1.insert(i)
        for i in self.init2:
            self.tree2.insert(i)
        self.assertTrue(contains_tree(self.tree1, self.tree2))

    def test_contains_tree_02(self):
        self.tree2.insert(95)
        self.assertFalse(contains_tree(self.tree1, self.tree2))


if __name__ == "__main__":
    unittest.main()

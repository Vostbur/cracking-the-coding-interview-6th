# Проверьте, пересекаются ли два заданных (одно-)связных списка.
# Верните узел пересечения. Учтите, что пересечение определяется ссылкой, а не значением.
# Иначе говоря, если k-й узел первого связного списка точно совпадает
# (по ссылке) с j-м узлом второго связного списка, то списки считаются пересекающимися
import unittest
from LinkedList import LinkedList


def intersection(l1, l2):
    if l1.last_node is not l2.last_node:
        return False

    diff = len(l1) - len(l2)
    if diff >= 0:
        long, short = l1, l2
    else:
        long, short = l2, l1

    node_short, node_long = short.start_node, long.start_node
    for i in range(abs(diff)):
        node_long = node_long.ref

    while node_short is not node_long:
        node_short, node_long = node_short.ref, node_long.ref

    return node_short


class Test(unittest.TestCase):
    def test_intersection_true(self):
        l1 = l2 = LinkedList()
        l1.insert_multiple(1, 2, 3)
        result = intersection(l1, l2)
        self.assertTrue(result)

    def test_intersection_false(self):
        l1 = LinkedList()
        l1.insert_multiple(1, 2, 3)
        l2 = LinkedList()
        l2.insert_multiple(4, 5, 6, 7)
        self.assertFalse(intersection(l1, l2))


if __name__ == '__main__':
    unittest.main()

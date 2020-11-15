# Реализуйте алгоритм для нахождения в односвязном списке k-го элемента с конца
import unittest
from LinkedList import LinkedList


def search_from_end(linked_list, index):
    i = linked_list.get_count() - index
    if i < 0:
        return None
    if i == 0:
        return linked_list.start_node.item
    n = linked_list.start_node
    for j in range(i):
        n = n.ref
    return n.item


class Test(unittest.TestCase):
    def test_delete_middle(self):
        ll = LinkedList()
        ll.insert_multiple(1, 2, 3, 4)
        self.assertEqual(search_from_end(ll, 2), 3)


if __name__ == "__main__":
    unittest.main()

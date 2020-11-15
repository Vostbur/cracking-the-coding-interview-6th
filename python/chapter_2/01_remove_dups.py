# Напишите код для удаления дубликатов из несортированного связного списка.
# Дополнительно:
# Как вы будете решать задачу, если использовать временный буфер запрещено?
import unittest
from LinkedList import LinkedList


def remove_dups(linked_list):
    if linked_list.start_node is None:
        return
    s = set([linked_list.start_node.item])
    n = linked_list.start_node
    while n.ref:
        if n.ref.item in s:
            n.ref = n.ref.ref
        else:
            s.add(n.ref.item)
            n = n.ref


def remove_dups_two_refs(linked_list):
    if linked_list.start_node is None:
        return
    n = linked_list.start_node
    while n:
        next_ref = n
        while next_ref.ref:
            if next_ref.ref.item == n.item:
                next_ref.ref = next_ref.ref.ref
            else:
                next_ref = next_ref.ref
        n = n.ref


class Test(unittest.TestCase):
    def test_remove_dups(self):
        ll = LinkedList()
        ll.insert_multiple(1, 2, 2, 3, 3, 3)
        remove_dups(ll)
        self.assertEqual(str(ll), "1,2,3")

    def test_remove_dups_two_refs(self):
        ll = LinkedList()
        ll.insert_multiple(1, 2, 2, 3, 3, 3)
        remove_dups_two_refs(ll)
        self.assertEqual(str(ll), "1,2,3")


if __name__ == "__main__":
    unittest.main()

import unittest
from LinkedList import LinkedList, LinkedListException


class TestLinkedList(unittest.TestCase):
    linked_list = LinkedList()
    linked_list_with_args = LinkedList(1, 2, 3)

    def test_01_repr_of_null_list(self):
        self.assertEqual(str(self.linked_list), "")

    def test_02_insert_at_start(self):
        self.linked_list.insert_at_start(1)
        self.linked_list.insert_at_start(2)
        self.linked_list.insert_at_start(3)
        self.assertEqual(str(self.linked_list), "3,2,1")

    def test_03_insert_at_end(self):
        self.linked_list.insert_at_end(0)
        self.assertEqual(str(self.linked_list), "3,2,1,0")

    def test_04_insert_after_item(self):
        self.linked_list.insert_after_item(2, 4)
        self.assertEqual(str(self.linked_list), "3,2,4,1,0")

    def test_05_insert_after_item_exception(self):
        self.assertRaises(
            LinkedListException, self.linked_list.insert_after_item, 5, 0)

    def test_06_insert_before_item(self):
        self.linked_list.insert_before_item(1, 5)
        self.assertEqual(str(self.linked_list), "3,2,4,5,1,0")

    def test_07_insert_before_item_exception(self):
        self.assertRaises(
            LinkedListException, self.linked_list.insert_before_item, 6, 0
        )

    def test_08_insert_at_index(self):
        self.linked_list.insert_at_index(1, 6)
        self.assertEqual(str(self.linked_list), "3,6,2,4,5,1,0")

    def test_09_insert_at_index_exception(self):
        self.assertRaises(
            LinkedListException, self.linked_list.insert_at_index, 10, 0)

    def test_10_get_count(self):
        self.assertEqual(self.linked_list.get_count(), 7)

    def test_11_search_item_true(self):
        self.assertTrue(self.linked_list.search_item(3))

    def test_12_search_item_false(self):
        self.assertFalse(self.linked_list.search_item(10))

    def test_13_repr_of_full_linked_list(self):
        self.assertEqual(str(self.linked_list_with_args), "1,2,3")

    def test_14_delete_at_start(self):
        self.linked_list.delete_at_start()
        self.assertEqual(str(self.linked_list), "6,2,4,5,1,0")

    def test_15_delete_at_start_exception(self):
        new_linked_list = LinkedList()
        self.assertRaises(LinkedListException, new_linked_list.delete_at_start)

    def test_16_delete_at_end(self):
        self.linked_list.delete_at_end()
        self.assertEqual(str(self.linked_list), "6,2,4,5,1")

    def test_17_delete_at_end_exception(self):
        new_linked_list = LinkedList()
        self.assertRaises(LinkedListException, new_linked_list.delete_at_end)

    def test_18_delete_element_by_value(self):
        self.linked_list.delete_element_by_value(4)
        self.assertEqual(str(self.linked_list), "6,2,5,1")

    def test_19_delete_element_by_value_exception(self):
        self.assertRaises(
            LinkedListException, self.linked_list.delete_element_by_value, 10
        )

    def test_20_reverse(self):
        self.linked_list.reverse()
        self.assertEqual(str(self.linked_list), "1,5,2,6")

    def test_21_iter(self):
        result = ",".join(str(i) for i in self.linked_list)
        self.assertEqual(result, "1,5,2,6")


if __name__ == "__main__":
    unittest.main()

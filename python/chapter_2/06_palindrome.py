# Реализуйте функцию, проверяющую, является ли связный список палиндромом.
import unittest
from LinkedList import LinkedList


def is_palindrome(data):
    ll = LinkedList()
    ll.insert_multiple(*data)
    return str(ll) == str(ll.reverse())


def is_palindrome_stack(data):
    ll = LinkedList()
    ll.insert_multiple(*data)

    left = right = ll.start_node
    stack = []
    while right and right.ref:
        stack.append(left.item)
        left = left.ref
        right = right.ref.ref

    if right:
        left = left.ref

    while left:
        if stack.pop() != left.item:
            return False
        left = left.ref

    return True


class Test(unittest.TestCase):
    cases = [
        ([1, 2, 3, 4, 5, 4, 3, 2, 1], True),
        ([1, 2, 3, 4, 5, 6, 3, 2, 1], False)
    ]

    def test_palindrome(self):
        for data, expect in self.cases:
            self.assertEqual(expect, is_palindrome(data))

    def test_palindrome_stack(self):
        for data, expect in self.cases:
            self.assertEqual(expect, is_palindrome_stack(data))


if __name__ == "__main__":
    unittest.main()

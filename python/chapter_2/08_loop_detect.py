# Для кольцевого связного списка реализуйте алгоритм, возвращающий начальный узел петли.
# Определение:
# Кольцевой связный список — это связный список, в котором указатель следующего
# узла ссылается на более ранний узел, образуя петлю.
# Пример:
# Ввод: A->B->C->D->E->C (предыдущий узел C)
# Вывод: C
import unittest
from LinkedList import LinkedList


def loop_detect(ll):
    fast = slow = ll.start_node

    while fast and fast.ref:
        slow = slow.ref
        fast = fast.ref.ref
        if fast is slow:
            break

    if fast is None or fast.ref is None:
        return None

    slow = ll.start_node
    while fast is not slow:
        slow = slow.ref
        fast = fast.ref

    return fast


class Test(unittest.TestCase):
    def test_loop_detect(self):
        ll = LinkedList()
        ll.insert_multiple(1, 2, 3, 4, 5)
        loop_node = ll.start_node.ref.ref
        ll.last_node.ref = loop_node
        result = loop_detect(ll)
        print(result.item)
        self.assertEqual(result, loop_node)


if __name__ == '__main__':
    unittest.main()

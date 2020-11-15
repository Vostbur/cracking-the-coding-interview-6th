# Напишите код для разбиения связного списка вокруг значения x, так чтобы
# все узлы, меньшие x, предшествовали узлам, большим или равным x. Если x
# содержится в списке, то значения x должны следовать строго после элементов,
# меньших x (см. далее). Элемент разбивки x может находиться где угодно
# в «правой части»; он не обязан располагаться между левой и правой частью.
# Пример:
# Ввод: 3->5->8->5->10->2->1 [значение разбивки = 5]
# Вывод: 3->1->2->10->5->5->8
import unittest
from LinkedList import LinkedList


def divide(linked_list, value):
    n = linked_list.last_node = linked_list.start_node
    while n:
        next_node = n.ref
        n.ref = None
        if n.item < value:
            n.ref = linked_list.start_node
            linked_list.start_node = n
        else:
            linked_list.last_node.ref = n
            linked_list.last_node = n
        n = next_node


class Test(unittest.TestCase):
    def test_divide(self):
        ll = LinkedList()
        ll.insert_multiple(3, 5, 8, 5, 10, 2, 1)
        divide(ll, 3)
        self.assertEqual("1,2,3,5,8,5,10", str(ll))


if __name__ == "__main__":
    unittest.main()

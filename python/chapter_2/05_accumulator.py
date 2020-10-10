# Два числа хранятся в виде связных списков, в которых каждый узел представляет один разряд.
# Все цифры хранятся в обратном порядке, при этом
# младший разряд (единицы) хранится в начале списка. Напишите функцию,
# которая суммирует два числа и возвращает результат в виде связного списка.
# Пример:
# Ввод: (7->1->6) + (5->9->2), то есть 617 + 295.
# Вывод: 2->1->9, то есть 912.
# Дополнительно
# Решите задачу, предполагая, что цифры записаны в прямом порядке.
# Ввод: (6->1->7) + (2->9->5), то есть 617 + 295.
# Вывод: 9->1->2, то есть 912.
import unittest
from LinkedList import LinkedList


# Сумма элементов записанных в обратном порядке
def accumulator_unusual(ll1, ll2):
    result = LinkedList()
    a, b = ll1.start_node, ll2.start_node
    shift = 0
    while a or b:
        c = shift
        if a:
            c += a.item
            a = a.ref
        if b:
            c += b.item
            b = b.ref
        result.insert_at_end(c % 10)
        shift = c // 10
    if shift:
        result.insert_at_end(shift)
    return result


# Сумма элементов записанных в прямом порядке
def accumulator_usual(ll1, ll2):
    len_ll1 = len(ll1.reverse())
    len_ll2 = len(ll2.reverse())
    diff = len_ll1 - len_ll2
    if diff > 0:
        ll2.insert_multiple(diff * (0,))
    elif diff > 0:
        ll1.insert_multiple(-diff * (0,))
    return accumulator_unusual(ll1, ll2).reverse()


# Сумма элементов записанных в прямом порядке (вариант 2)
def accumulator_usual_2(ll1, ll2):
    diff = len(ll1) - len(ll2)
    if diff > 0:
        for i in range(diff):
            ll2.insert_at_start(0)
    elif diff < 0:
        for i in range(-diff):
            ll1.insert_at_start(0)

    a, b = ll1.start_node, ll2.start_node
    c = 0
    while a and b:
        c = c * 10 + a.item + b.item
        a = a.ref
        b = b.ref

    result = LinkedList()
    for i in str(c):
        result.insert_at_end(int(i))
    return result


class Test(unittest.TestCase):
    def test_accumulator_unusual(self):
        ll1 = LinkedList()
        ll1.insert_multiple(7, 1, 6)
        ll2 = LinkedList()
        ll2.insert_multiple(5, 9, 2)
        actual = accumulator_unusual(ll1, ll2)
        self.assertEqual('2,1,9', str(actual))

    def test_accumulator_usual(self):
        ll1 = LinkedList()
        ll1.insert_multiple(6, 1, 7)
        ll2 = LinkedList()
        ll2.insert_multiple(2, 9, 5)
        actual = accumulator_usual(ll1, ll2)
        self.assertEqual('9,1,2', str(actual))

    def test_accumulator_usual_2(self):
        ll1 = LinkedList()
        ll1.insert_multiple(6, 1, 7)
        ll2 = LinkedList()
        ll2.insert_multiple(2, 9, 5)
        actual = accumulator_usual_2(ll1, ll2)
        self.assertEqual('9,1,2', str(actual))


if __name__ == '__main__':
    unittest.main()

# Напишие функцию, которая проверяет, является ли заданная строка
# перестановкой палиндрома. (Палиндром - слово или фраза, одинаково
# читающиеся в прямом и обратном направлении;
# перестановка - строка, содержащая те же символы в другом порядке.)
# Палиндром не ограничивается словами из словаря.
# Пример: Tact Coa
# Вывод: True (перестановки: 'taco cat', 'atco cta' и т.д.)
import unittest
import string
from collections import Counter


def check_palindrome(s):
    p = [i for i in s.lower() if i in string.ascii_lowercase]
    length = len(p)
    if length < 3:
        return False

    c = Counter()
    for i in p:
        c[i] += 1

    # all symbols are the same
    if len(list(c)) == 1:
        return False

    flag = False
    for i in p:
        if c[i] % 2:
            if flag:
                return False
            flag = True
    return flag == (length % 2)


class Test(unittest.TestCase):
    data = [
        ("Tact Coa", True),
        ("jhsabckuj ahjsbckj", True),
        ("Able was I ere I saw Elba", True),
        ("So patient a nurse to nurse a patient so", False),
        ("Random Words", False),
        ("Not a Palindrome", False),
        ("no x in nixon", True),
        ("azAZ", True),
    ]

    def test_check_palindrome(self):
        for [s, result] in self.data:
            self.assertEqual(check_palindrome(s), result)


if __name__ == "__main__":
    unittest.main()

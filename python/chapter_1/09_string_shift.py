"""Допустим, что существует метод isSubstring, проверяющий,
является ли одно слово подстрокой другого. Для двух строк s1 и s2 напишите код,
который проверяет, получена ли строка s2 циклическим сдвигом s1,
используя только один вызов метода isSubstring
(пример: слово waterbottle получено циклическим сдвигом erbottlewat)."""
import unittest


def isSubstring(s, sub):
    return s.find(sub) != -1


def is_shift(s, shift):
    return len(s) == len(shift) and isSubstring(s * 2, shift)


class Test(unittest.TestCase):
    data = [
        ("waterbottle", "erbottlewat", True),
        ("waterbottle", "elttobretaw", False),
        ("foo", "bar", False),
        ("foo", "foofoo", False),
    ]

    def test_is_shift(self):
        for [s, shift, result] in self.data:
            self.assertEqual(is_shift(s, shift), result)


if __name__ == "__main__":
    unittest.main()

"""Существуют три вида модифицирующих операций над строками:
вставка символа, удаление символа и замена символа. Напишите
функцию, которая проверяет, находятся ли две строки на расстоянии
одной модификации (или нуля модификаций).
Пример:
pale, ple -> true
pales, pale -> true
pale, bale -> true
pale, bake -> false"""
import unittest


def check_modification(s1, s2):
    if s1 == s2:
        return True
    if len(s1) == len(s2):
        for i in range(1, len(s1) + 1):
            if s1[: i - 1] + s1[i:] == s2[: i - 1] + s2[i:]:
                return True
    else:
        max_s = max(s1, s2, key=len)
        min_s = s2 if max_s == s1 else s1
        for i in range(1, len(max_s) + 1):
            if max_s[: i - 1] + max_s[i:] == min_s:
                return True
    return False


class Test(unittest.TestCase):
    data = [
        ("pale", "ple", True),
        ("pales", "pale", True),
        ("pale", "bale", True),
        ("paleabc", "pleabc", True),
        ("pale", "ble", False),
        ("a", "b", True),
        ("", "d", True),
        ("d", "de", True),
        ("pale", "pale", True),
        ("pale", "ple", True),
        ("ple", "pale", True),
        ("pale", "bale", True),
        ("pale", "bake", False),
        ("pale", "pse", False),
        ("ples", "pales", True),
        ("pale", "pas", False),
        ("pas", "pale", False),
        ("pale", "pkle", True),
        ("pkle", "pable", False),
        ("pal", "palks", False),
        ("palks", "pal", False),
    ]

    def test_func(self):
        for [s1, s2, result] in self.data:
            self.assertEqual(check_modification(s1, s2), result)


if __name__ == "__main__":
    unittest.main()

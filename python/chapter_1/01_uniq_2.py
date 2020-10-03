"""Реализуйте алгоритм, определяющий, все ли символы в строке встречаются только один раз.
А если при этом запрещено использование дополнительных структур данных?
"""
import unittest


def unique(s):
    if len(s) < 2:
        return True
    d = [s[0]]
    for i in s[1:]:
        if i in d:
            return False
        d.append(i)
    return True


class Test(unittest.TestCase):
    dataset_true = ["abcd 1234.:", "", "Whats'up"]
    dataset_false = ["They will", "aasdf", "asdfgg"]

    def test_unique(self):
        for s in self.dataset_true:
            result = unique(s)
            self.assertTrue(result)

        for s in self.dataset_false:
            result = unique(s)
            self.assertFalse(result)


if __name__ == "__main__":
    unittest.main()

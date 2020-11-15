# Для двух строк напишите метод, определяющий,
# является ли одна строка перестановкой другой
import unittest
from collections import Counter


def permutation(s1, s2):
    # O(N)
    if s1 == s2 or len(s1) != len(s2):
        return False
    c = Counter()
    for i in s1:
        c[i] += 1
    for i in s2:
        if c[i] == 0:
            return False
        c[i] -= 1
    return True


class Test(unittest.TestCase):
    dataset_false = [("asd", "asd"), ("asd", "asdf"), ("", "")]
    dataset_true = [("А вас", "Авас "), ("asd", "sda")]

    def test_permutation(self):
        # false check
        for i in self.dataset_false:
            result = permutation(*i)
            self.assertFalse(result)
        # true check
        for i in self.dataset_true:
            result = permutation(*i)
            self.assertTrue(result)


if __name__ == "__main__":
    unittest.main()

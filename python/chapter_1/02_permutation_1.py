"""Для двух строк напишите метод, определяющий, является ли одна строка перестановкой другой"""
import unittest


def permutation(s1, s2):
    if s1 == s2 or len(s1) != len(s2):
        return False
    s1, s2 = sorted(s1), sorted(s2)
    return True if s1 == s2 else False


class Test(unittest.TestCase):
    dataset_false = [('asd', 'asd'), ('asd', 'asdf'), ('', '')]
    dataset_true = [('А вас', 'Авас '), ('asd', 'sda')]

    def test_permutation(self):
        # False check
        for i in self.dataset_false:
            result = permutation(*i)
            self.assertFalse(result)
        # True check
        for i in self.dataset_true:
            result = permutation(*i)
            self.assertTrue(result)


if __name__ == "__main__":
    unittest.main()

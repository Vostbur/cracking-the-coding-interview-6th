# Для двух строк напишите метод, определяющий,
# является ли одна строка перестановкой другой
import unittest


def permutation(s1, s2):
    if s1 == s2 or len(s1) != len(s2):
        return False
    d = {}
    for i in s1:
        d[i] = d[i] + 1 if i in d else 1
    for i in s2:
        if not (i in d):
            return False
        if d[i] != 0:
            d[i] -= 1
        else:
            return False
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

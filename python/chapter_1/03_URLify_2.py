"""Напишите метод, заменяющий все пробелы в строке символами '%20'.
Можете считать, что длина строки позволяет сохранить дополнительные символы,
а фактическая длина строки известна заранее.
Пример:
Ввод: 'Mr John Smith    ', 13
Вывод: 'Mr%20John%20Smith"""
import unittest


def URLify(s, l):
    # O(N)
    ind = len(s)

    for i in reversed(range(l)):
        if s[i] == ' ':
            s[ind - 3:ind] = '%20'
            ind -= 3
        else:
            s[ind - 1] = s[i]
            ind -= 1

    return s


class Test(unittest.TestCase):
    dataset = [(list('as d  '), 4, list('as%20d')),
               (list('Mr John Smith    '), 13, list('Mr%20John%20Smith'))]

    def test_URLify(self):
        for [test_string, length, result] in self.dataset:
            self.assertEqual(URLify(test_string, length), result)


if __name__ == '__main__':
    unittest.main()

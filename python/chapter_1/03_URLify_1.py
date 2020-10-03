"""Напишите метод, заменяющий все пробелы в строке символами '%20'"""
import unittest


def URLify(s):
    # O(N)
    url = s.rstrip()
    return url.replace(' ', r'%20', url.count(' '))


class Test(unittest.TestCase):
    dataset = {'as d': r'as%20d', 'Mr John Smith         ': r'Mr%20John%20Smith'}

    def test_URLify(self):
        for k, v in self.dataset.items():
            self.assertEqual(v, URLify(k))


if __name__ == '__main__':
    unittest.main()

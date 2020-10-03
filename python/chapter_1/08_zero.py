"""Напишите алгоритм, реализующий условие:
если элемент матрицы MxN равен 0,
то весь столбец и вся строка обнуляются"""
import unittest


def zero(m):
    row = len(m)
    col = len(m[0])
    z = []
    for i in range(row):
        for j in range(col):
            if m[i][j] == 0:
                z.append([i, j])
    for [r, c] in z:
        for i in range(col):
            m[r][i] = 0
        for i in range(row):
            m[i][c] = 0
    return m


class Test(unittest.TestCase):
    data = [([[1, 2, 3, 4, 0],
              [6, 0, 8, 9, 10],
              [11, 12, 13, 14, 15],
              [16, 0, 18, 19, 20],
              [21, 22, 23, 24, 25]],
             [[0, 0, 0, 0, 0],
              [0, 0, 0, 0, 0],
              [11, 0, 13, 14, 0],
              [0, 0, 0, 0, 0],
              [21, 0, 23, 24, 0]])]

    def test_zero(self):
        for [m, result] in self.data:
            self.assertEqual(zero(m), result)


if __name__ == '__main__':
    unittest.main()

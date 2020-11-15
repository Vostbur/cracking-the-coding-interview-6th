"""Имеется изображение, представленное матрицей NxN;
каждый пиксел представлен 4 байтами.
Напишите метод для поворота изображения на 90 градусов.
Удастся ли вам выполнить эту операцию 'на месте'?"""
import unittest


def rotate(m):
    n = len(m)
    for row in range(n // 2):
        start, end = row, n - row - 1
        for col in range(start, end):
            element = m[row][col]
            m[row][col] = m[-col - 1][row]
            m[-col - 1][row] = m[-row - 1][-col - 1]
            m[-row - 1][-col - 1] = m[col][-row - 1]
            m[col][-row - 1] = element
    return m


class Test(unittest.TestCase):
    data = [
        (
            [
                [1, 2, 3, 4, 5],
                [6, 7, 8, 9, 10],
                [11, 12, 13, 14, 15],
                [16, 17, 18, 19, 20],
                [21, 22, 23, 24, 25],
            ],
            [
                [21, 16, 11, 6, 1],
                [22, 17, 12, 7, 2],
                [23, 18, 13, 8, 3],
                [24, 19, 14, 9, 4],
                [25, 20, 15, 10, 5],
            ],
        )
    ]

    def test_rotate(self):
        for [m, result] in self.data:
            self.assertEqual(rotate(m), result)


if __name__ == "__main__":
    unittest.main()

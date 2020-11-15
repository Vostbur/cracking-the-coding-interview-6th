# Реализуйте метод для выполнения простейшего сжатия строк с использованием
# счетчика повторяющихся символов. Например, строка aabcccccaa превращается
# в a2b1c5a3. Если 'сжатая' строка не становится короче исходной, то метод
# возвращает исходную строку. Предполагается, что строка состоит только из
# букв верхнего и нижнего регистра (a-z)
import unittest


def compression(s):
    length = len(s)
    if length < 3:
        return s

    count, c = 1, []
    for i in range(1, length):
        if s[i] != s[i - 1]:
            c.append(s[i - 1] + str(count))
            count = 0
        count += 1
    c.append(s[-1] + str(count))
    return min(s, "".join(c), key=len)


class Test(unittest.TestCase):
    data = [
        ("aabcccccaaa", "a2b1c5a3"), ("abbbbbb", "a1b6"), ("abcdef", "abcdef")
    ]

    def test_compression(self):
        for [s, result] in self.data:
            self.assertEqual(compression(s), result)


if __name__ == "__main__":
    unittest.main()

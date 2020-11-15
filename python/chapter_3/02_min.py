# Как реализовать стек, в котором кроме стандартных функций push и pop будет
# поддерживаться функция min, возвращающая минимальный элемент? Все
# операции — push, pop и min — должны выполняться за время O(1).
import unittest


class Stack:
    def __init__(self):
        self.stack = []
        self.min_values = []
        self.length = 0

    def push(self, x):
        m = x if self.is_empty() else self.min()
        if x < m:
            m = x
        self.min_values.append(m)
        self.stack.append(x)
        self.length += 1

    def pop(self):
        if self.is_empty():
            return Exception("Cannot return pop. Stack is empty.")
        ret = self.stack[self.length - 1]
        self.stack = self.stack[: self.length - 1]
        self.min_values = self.min_values[: self.length - 1]
        self.length -= 1
        return ret

    def min(self):
        if self.is_empty():
            return Exception("Cannot return min. Stack is empty.")
        return self.min_values[self.length - 1]

    def is_empty(self):
        return self.length == 0


class Test(unittest.TestCase):
    s = Stack()

    def test_Stack1(self):
        self.s.push(10)
        self.s.push(11)
        self.s.push(9)
        self.assertEqual(self.s.min(), 9)

    def test_Stack2(self):
        self.s.pop()
        self.assertEqual(self.s.min(), 10)


if __name__ == "__main__":
    unittest.main()

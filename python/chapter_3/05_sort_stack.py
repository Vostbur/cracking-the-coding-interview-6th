# Напишите программу сортировки стека, в результате которой наименьший
# элемент оказывается на вершине стека. Вы можете использовать дополнительный временный стек,
# но элементы не должны копироваться в другие
# структуры данных (например, в массив). Стек должен поддерживать следующие операции: push, pop, peek, isEmpty.
import unittest
from stack import Stack, StackException


class SortStack:
    def __init__(self):
        self.sort_stack = Stack()
        self.temp_stack = Stack()

    def __str__(self):
        return f'Sorted stack: {str(self.sort_stack)}Temp stack: {str(self.temp_stack)}'

    def __iter__(self):
        return self.sort_stack.__iter__()

    def push(self, x):
        # pop values until empty or fine the larger one
        if not self.sort_stack.is_empty():
            while True:
                try:
                    peek = self.sort_stack.peek()
                except StackException:
                    break
                if peek < x:
                    pop = self.sort_stack.pop()
                    self.temp_stack.push(pop)
                else:
                    break

        self.sort_stack.push(x)
        # put back smaller values
        if not self.temp_stack.is_empty():
            for i in self.temp_stack:
                self.sort_stack.push(i)

    def pop(self):
        if self.sort_stack.is_empty():
            raise StackException('Cannot pop. Stack is empty')
        return self.sort_stack.pop()

    def peek(self):
        if self.sort_stack.is_empty():
            raise StackException('Cannot peek. Stack is empty')
        return self.sort_stack.peek()

    def is_empty(self):
        return self.sort_stack.is_empty()


class Test(unittest.TestCase):
    ss = SortStack()

    def test_01_SortStack(self):
        init = [1, 3, 6, 5, 7, 2, 4]
        expected = [1, 2, 3, 4, 5, 6, 7]

        for i in init:
            self.ss.push(i)
        actual = []
        print(self.ss)
        for i in self.ss:
            actual.append(i)
        self.assertEqual(actual, expected)


if __name__ == '__main__':
    unittest.main()

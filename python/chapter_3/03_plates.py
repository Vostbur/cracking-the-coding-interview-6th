# Как известно, слишком высокая стопка тарелок может развалиться.
# Следовательно, в реальной жизни, когда высота стопки превысила бы некоторое
# пороговое значение, мы начали бы складывать тарелки в новую стопку.
# Реализуйте структуру данных SetOfStacks, имитирующую реальную ситуацию.
# Структура SetOfStack должна состоять из нескольких стеков, новый стек
# создается, как только предыдущий достигнет порогового значения. Методы
# SetOfStacks.push() и SetOfStacks.pop() должны вести себя так же, как при
# работе с одним стеком (то есть метод pop() должен возвращать те же значения,
# которые бы он возвращал при использовании одного большого стека).
# Дополнительно:
# Реализуйте функцию popAt(int index), которая осуществляет операцию pop
# с заданным внутренним стеком.
import unittest


class Node:
    def __init__(self, value):
        self.value = value
        self.above = None
        self.below = None


class Stack:
    def __init__(self, capacity):
        self.size = 0
        self.top = None
        self.bottom = None
        self.capacity = capacity

    def __iter__(self):
        n = self.top
        while n:
            v = n.value
            n = n.below
            yield v

    def __str__(self):
        return ','.join([str(i) for i in self])

    def push(self, x):
        new_node = Node(x)
        self.size += 1
        if self.size == 1:
            self.bottom = self.top = new_node
            return
        self.top.above = new_node
        new_node.below = self.top
        self.top = new_node

    def pop(self):
        if not self.top:
            return None
        n = self.top
        self.top = self.top.below
        self.size -= 1
        return n.value

    def from_bottom(self):
        n = self.bottom
        self.bottom = self.bottom.above
        if self.bottom:
            self.bottom.below = None
        self.size -= 1
        return n.value

    def is_full(self):
        return self.size == self.capacity


class StackPlates:
    def __init__(self, capacity):
        self.capacity = capacity
        self.stacks = []

    def push(self, x):
        stack = self.get_last_stack()
        if stack and not stack.is_full():
            stack.push(x)
        else:
            new_stack = Stack(self.capacity)
            new_stack.push(x)
            self.stacks.append(new_stack)

    def pop(self):
        stack = self.get_last_stack()
        if not stack:
            return None
        v = stack.pop()
        if stack.size == 0:
            del self.stacks[-1]
        return v

    def pop_at(self, index):
        return self.shift(index, True)

    def shift(self, index, from_top):
        stack = self.stacks[index]
        item = stack.pop() if from_top else stack.from_bottom()
        if stack.size == 0:
            del self.stacks[index]
        elif len(self.stacks) > index + 1:
            v = self.shift(index + 1, False)
            stack.push(v)
        return item

    def get_last_stack(self):
        if not self.stacks:
            return None
        return self.stacks[-1]


class Test(unittest.TestCase):
    s = StackPlates(5)

    def test_01_StackPlates(self):
        for i in range(15):
            self.s.push(i)
        actual = []
        for _ in range(15):
            actual.append(self.s.pop())
        self.assertEqual(list(reversed(range(15))), actual)

    def test_02_pop_at(self):
        for i in range(15):
            self.s.push(i)
        actual = []
        for _ in range(11):
            actual.append(self.s.pop_at(0))
        self.assertEqual(list(range(4, 15)), actual)


if __name__ == '__main__':
    unittest.main()

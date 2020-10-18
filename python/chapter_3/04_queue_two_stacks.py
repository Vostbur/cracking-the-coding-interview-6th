# Напишите класс MyQueue, который реализует очередь с использованием двух стеков.
import unittest
from stack import Stack, StackException


class Queue:
    def __init__(self):
        self.master = Stack()
        self.slave = Stack()

    def push(self, x):
        self.master.push(x)

    def pop(self):
        for i in self.master:
            self.slave.push(i)
        ret = self.slave.pop()
        for i in self.slave:
            self.master.push(i)
        return ret


class Test(unittest.TestCase):
    stack = Stack()
    queue = Queue()

    def test_01_push(self):
        self.stack.push(1)
        self.stack.push(2)
        self.stack.push(3)
        self.assertEqual(str(self.stack), 'stack:[1, 2, 3] length:3\n')

    def test_02_pop(self):
        self.stack.pop()
        self.assertEqual(str(self.stack), 'stack:[1, 2] length:2\n')

    def test_03_pop(self):
        self.stack.pop()
        self.stack.pop()
        with self.assertRaises(StackException) as e:
            self.stack.pop()
        exc = e.exception
        self.assertEqual(str(exc), 'Cannot pop. Stack is empty')

    def test_04_queue(self):
        self.queue.push(1)
        self.queue.push(2)
        self.queue.push(3)
        actual = [self.queue.pop(), self.queue.pop(), self.queue.pop()]
        self.assertEqual(actual, [1, 2, 3])


if __name__ == '__main__':
    unittest.main()

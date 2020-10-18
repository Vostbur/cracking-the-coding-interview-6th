class StackException(Exception):
    pass


class Stack:
    def __init__(self):
        self.stack = []
        self.length = 0

    def __str__(self):
        return f'stack:{self.stack} length:{self.length}\n'

    def __len__(self):
        return self.length

    def __iter__(self):
        while self.length:
            yield self.pop()

    def push(self, x):
        self.stack.append(x)
        self.length += 1

    def pop(self):
        if self.is_empty():
            raise StackException('Cannot pop. Stack is empty')
        l = self.length - 1
        r = self.stack[l]
        self.stack = self.stack[:l]
        self.length -= 1
        return r

    def peek(self):
        if self.is_empty():
            raise StackException('Cannot peek. Stack is empty')
        return self.stack[self.length - 1]

    def is_empty(self):
        return self.length == 0

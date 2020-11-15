# Опишите, как бы вы использовали один одномерный массив для реализации
# трех стеков
class MultiStack:
    def __init__(self, stack_size):
        self.number_of_stacks = 3
        self.stack = [0] * (self.number_of_stacks * stack_size)
        self.sizes = [0] * self.number_of_stacks
        self.stack_size = stack_size

    def __str__(self):
        return (
            f"number_of_stacks={self.number_of_stacks}\n"
            f"stack={self.stack}\n"
            f"sizes={self.sizes}\n"
            f"stack_size={self.stack_size}\n"
        )

    def push(self, stack_num, value):
        if self.is_full(stack_num):
            raise Exception("Stack is full")
        self.sizes[stack_num] += 1
        self.stack[self.index(stack_num)] = value

    def pop(self, stack_num):
        if self.is_empty(stack_num):
            raise Exception("Stack is empty")
        ind = self.index(stack_num)
        ret, self.stack[ind] = self.stack[ind], 0
        self.sizes[stack_num] -= 1
        return ret

    def peek(self, stack_num):
        if self.is_empty(stack_num):
            raise Exception("Stack is empty")
        return self.stack[self.index(stack_num)]

    def is_empty(self, stack_num):
        return self.sizes[stack_num] == 0

    def is_full(self, stack_num):
        return self.sizes[stack_num] == self.stack_size

    def index(self, stack_num):
        return stack_num * self.stack_size + self.sizes[stack_num] - 1


if __name__ == "__main__":
    s = MultiStack(2)
    s.push(0, 1)
    s.push(0, 2)
    print(s)
    try:
        s.push(0, 3)
    except Exception as e:
        print(e)
    print(s.is_empty(1))
    s.push(1, 10)
    s.push(1, 20)
    s.push(2, 100)
    s.push(2, 200)
    print(s)
    print(s.peek(2))
    print(s.pop(1))
    print(s.pop(1))
    try:
        s.pop(1)
    except Exception as e:
        print(e)
    print(s.is_empty(1))
    print(s)

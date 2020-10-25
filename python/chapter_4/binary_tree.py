class BinaryTree:
    def __init__(self, root):
        self.value = root
        self.size = 1
        self.left = None
        self.right = None
        self.parent = None

    def __len__(self):
        return self.size

    def __str__(self):
        return self.print(self, 0, '', '')

    def insert(self, value):
        if value <= self.value:
            if not self.left:
                self.setLeftChild(BinaryTree(value))
            else:
                self.left.insert(value)
        else:
            if not self.right:
                self.setRightChild(BinaryTree(value))
            else:
                self.right.insert(value)
        self.size += 1

    def find(self, value):
        if value == self.value:
            return self
        else:
            if value <= self.value:
                return self.left if self.left.find(value) else None
            elif value > self.value:
                return self.right if self.right.find(value) else None
        return None

    def setLeftChild(self, left):
        self.left = left
        self.left.parent = self

    def setRightChild(self, right):
        self.right = right
        self.right.parent = self

    def print(self, n, level, ch, output):
        if n:
            output += self.print(n.left, level + 1, '┌', '') + \
                      '  ' * level * 4 + ch + '----->' + str(n.value) + '\n' + \
                      self.print(n.right, level + 1, '└', '')
        return output

from random import randint


class LinkedListException(Exception):
    pass


class Node:
    def __init__(self, data):
        self.item = data
        self.ref = None

    def __str__(self):
        return str(self.item)


class LinkedList:
    def __init__(self, *args):
        self.start_node = None
        self.last_node = None
        for arg in args:
            self.insert_at_end(arg)

    def __str__(self):
        if self.start_node is None:
            return ""
        elements = []
        n = self.start_node
        while n:
            elements.append(str(n.item))
            n = n.ref
        return ",".join(elements)

    def __len__(self):
        return self.get_count()

    def __iter__(self):
        n = self.start_node
        while n:
            yield n
            n = n.ref

    def insert_at_start(self, data):
        new_node = Node(data)
        new_node.ref = self.start_node
        self.start_node = new_node

    def insert_at_end(self, data):
        new_node = Node(data)
        if self.start_node is None:
            self.start_node = new_node
            return self.start_node
        n = self.start_node
        while n.ref:
            n = n.ref
        n.ref = new_node
        self.last_node = n.ref
        return self.last_node

    def insert_after_item(self, x, data):
        n = self.start_node
        while n:
            if n.item == x:
                break
            n = n.ref
        if n is None:
            raise LinkedListException("item not in the list")
        new_node = Node(data)
        new_node.ref = n.ref
        n.ref = new_node

    def insert_before_item(self, x, data):
        if self.start_node is None:
            raise LinkedListException("list has no element")
        if x == self.start_node.item:
            new_node = Node(data)
            new_node.ref = self.start_node
            self.start_node = new_node
            return
        n = self.start_node
        while n.ref:
            if n.ref.item == x:
                break
            n = n.ref
        if n.ref is None:
            raise LinkedListException("item not in the list")
        new_node = Node(data)
        new_node.ref = n.ref
        n.ref = new_node

    def insert_at_index(self, index, data):
        if index == 0:
            self.insert_at_start(data)
            return
        i = 1
        n = self.start_node
        while i < index - 1 and n:
            n = n.ref
            i += 1
        if n is None:
            raise LinkedListException("index out of bound")
        new_node = Node(data)
        new_node.ref = n.ref
        n.ref = new_node

    def insert_multiple(self, *args):
        for arg in args:
            self.insert_at_end(arg)

    def get_count(self):
        if self.start_node is None:
            return 0
        n = self.start_node
        count = 0
        while n:
            count += 1
            n = n.ref
        return count

    def search_item(self, x):
        if self.start_node is None:
            raise LinkedListException("list has no elements")
        n = self.start_node
        while n:
            if n.item == x:
                return True
            n = n.ref
        return False

    def delete_at_start(self):
        if self.start_node is None:
            raise LinkedListException("the list has no element to delete")
        self.start_node = self.start_node.ref

    def delete_at_end(self):
        if self.start_node is None:
            raise LinkedListException("the list has no element to delete")
        n = self.start_node
        while n.ref.ref:
            n = n.ref
        n.ref = None

    def delete_element_by_value(self, x):
        if self.start_node is None:
            raise LinkedListException("the list has no element to delete")
        if self.start_node.item == x:
            self.start_node = self.start_node.ref
            return
        n = self.start_node
        while n.ref:
            if n.ref.item == x:
                break
            n = n.ref
        if n.ref is None:
            raise LinkedListException("item not found in the list")
        n.ref = n.ref.ref

    def reverse(self):
        prev_node = None
        n = self.start_node
        while n:
            next_node = n.ref
            n.ref = prev_node
            prev_node = n
            n = next_node
        self.start_node = prev_node
        return self

    def generate(self, x):
        for i in range(x):
            self.insert_at_end(randint(0, 100))


if __name__ == "__main__":
    pass

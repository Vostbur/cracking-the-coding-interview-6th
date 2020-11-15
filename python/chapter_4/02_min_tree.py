# Напишите алгоритм создания бинарного дерева поиска с минимальной
# высотой для отсортированного (по возрастанию) массива с уникальными
# целочисленными элементами.
class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


def mit_tree_from_sorted_list(array):
    return min_tree(array, 0, len(array) - 1)


def min_tree(array, start, end):
    if start > end:
        return
    mid = int((start + end) / 2)
    n = Node(array[mid])
    n.left = min_tree(array, start, mid - 1)
    n.right = min_tree(array, mid + 1, end)
    return n


def print_tree(n, level, ch):
    if n:
        print_tree(n.left, level + 1, "┌")
        print("  " * level * 4 + ch + "----->", n.value)
        print_tree(n.right, level + 1, "└")


if __name__ == "__main__":
    array = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 111, 1000]
    print_tree(mit_tree_from_sorted_list(array), 0, "")

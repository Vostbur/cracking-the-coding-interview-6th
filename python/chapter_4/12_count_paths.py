# Дано бинарное дерево, в котором каждый узел содержит целое число
# (положительное или отрицательное). Разработайте алгоритм для подсчета
# всех путей, сумма значений которых соответствует заданной величине.
# Обратите внимание, что путь не обязан начинаться или заканчиваться
# в корневом или листовом узле, но он должен идти вниз (переходя только
# от родительских узлов к дочерним).
import unittest
from chapter_4.binary_tree import BinaryTree


def count_path(root: BinaryTree, target_sum: int) -> int:
    return count_path_with_sum(root, target_sum, 0, dict())


def count_path_with_sum(
    node: BinaryTree, target_sum: int, running_sum: int, path_count: dict
) -> int:
    if not node:
        return 0

    running_sum += node.value
    summa = running_sum - target_sum
    total_paths = path_count.get(summa, 0)

    if running_sum == target_sum:
        total_paths += 1

    inc_dict(path_count, running_sum, 1)
    total_paths += count_path_with_sum(
        node.left, target_sum, running_sum, path_count)
    total_paths += count_path_with_sum(
        node.right, target_sum, running_sum, path_count)
    inc_dict(path_count, running_sum, -1)

    return total_paths


def inc_dict(hash_table: dict, key: int, delta: int) -> None:
    new_count = hash_table.get(key, 0) + delta
    if not new_count:
        hash_table.pop(key)
    else:
        hash_table[key] = new_count


class Test(unittest.TestCase):
    def test_count_paths(self):
        tree = BinaryTree(50)
        tree.insert(-20)
        tree.insert(-10)
        tree.insert(-10)
        tree.insert(-20)
        tree.insert(20)
        tree.insert(55)
        tree.insert(8)
        tree.insert(15)
        tree.insert(5)
        tree.insert(-10)
        print(tree)
        self.assertEqual(3, count_path(tree, 10))


if __name__ == "__main__":
    unittest.main()

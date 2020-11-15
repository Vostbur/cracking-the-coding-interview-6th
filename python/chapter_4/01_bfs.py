# Для заданного направленного графа разработайте алгоритм, проверяющий
# существование маршрута между двумя узлами.
import collections
import unittest


# used 'breadth first search' algorithm
def bfs(graph, start, end):
    visited, queue = set(), collections.deque([start])
    visited.add(start)
    while queue:
        node = queue.popleft()
        for neighbour in graph[node]:
            if neighbour not in visited:
                if neighbour == end:
                    return True
                visited.add(neighbour)
                queue.append(neighbour)
    return False


class Test(unittest.TestCase):
    graph = {0: [1, 2], 1: [2], 2: [3], 3: [1, 2]}

    def test_bfs_true(self):
        self.assertTrue(bfs(self.graph, 0, 2))

    def test_bfs_false(self):
        self.assertFalse(bfs(self.graph, 1, 4))


if __name__ == "__main__":
    unittest.main()

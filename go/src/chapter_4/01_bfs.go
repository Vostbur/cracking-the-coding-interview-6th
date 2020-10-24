// Для заданного направленного графа разработайте алгоритм, проверяющий
// существование маршрута между двумя узлами.
package chapter_4

// used 'breadth first search' algorithm
func bfs(graph map[int][]int, start, end int) bool {
	visited, queue := []int{}, GetQueue()
	visited = append(visited, start)
	queue.Add(start)

	for {
		node, err := queue.Remove()
		if err != nil {
			break
		}
		for _, neighbour := range graph[node] {
			if ok := find(visited, neighbour); !ok {
				if neighbour == end {
					return true
				}
				visited = append(visited, neighbour)
				queue.Add(neighbour)
			}
		}
	}
	return false
}

func find(v []int, n int) bool {
	for _, i := range v {
		if i == n {
			return true
		}
	}
	return false
}

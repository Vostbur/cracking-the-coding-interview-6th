package chapter_4

import "testing"

func TestBfs(t *testing.T) {
	graph := map[int][]int{0: {1, 2}, 1: {2}, 2: {3}, 3: {1, 2}}
	if !bfs(graph, 0, 3) {
		t.Errorf("Expected: true. Actual: false")
	}
	if bfs(graph, 0, 4) {
		t.Errorf("Expected: false. Actual: true")
	}
}

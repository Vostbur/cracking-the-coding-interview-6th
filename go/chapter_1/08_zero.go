// Напишите алгоритм, реализующий условие:
// если элемент матрицы MxN равен 0,
// то весь столбец и вся строка обнуляются
package chapter_1

func Zero(matrix [][]int) {
	var tmp [][2]int

	m, n := len(matrix), len(matrix[0])
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 0 {
				tmp = append(tmp, [2]int{row, col})
			}
		}
	}

	for _, pos := range tmp {
		for i := 0; i < n; i++ {
			matrix[pos[0]][i] = 0
		}
		for i := 0; i < m; i++ {
			matrix[i][pos[1]] = 0
		}
	}
}

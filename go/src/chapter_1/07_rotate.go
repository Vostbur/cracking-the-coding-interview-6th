// Имеется изображение, представленное матрицей NxN;
// каждый пиксел представлен 4 байтами.
// Напишите метод для поворота изображения на 90 градусов.
// Удастся ли вам выполнить эту операцию 'на месте'?
package chapter_1

func Rotate(matrix [][]int) {
	n := len(matrix)

	for row := 0; row < n/2; row++ {
		end := n - row - 1
		for i := row; i < end; i++ {
			offset := i - row
			tmp := matrix[row][i]
			matrix[row][i] = matrix[end-offset][row]
			matrix[end-offset][row] = matrix[end][end-offset]
			matrix[end][end-offset] = matrix[i][end]
			matrix[i][end] = tmp
		}
	}
}

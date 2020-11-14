// Дано бинарное дерево, в котором каждый узел содержит целое число
// (положительное или отрицательное). Разработайте алгоритм для подсчета
// всех путей, сумма значений которых соответствует заданной величине.
// Обратите внимание, что путь не обязан начинаться или заканчиваться
// в корневом или листовом узле, но он должен идти вниз (переходя только
// от родительских узлов к дочерним).
package chapter_4

func countPaths(root *BinaryTree, targetSum int) int {
	return countPathsWithSum(root, targetSum, 0, make(map[int]int))
}

func countPathsWithSum(node *BinaryTree, targetSum int, runningSum int, pathCount map[int]int) int {
	if node == nil {
		return 0
	}

	runningSum += node.value
	sum := runningSum - targetSum
	totalPaths := pathCount[sum]

	if runningSum == targetSum {
		totalPaths++
	}

	incrementHashTable(pathCount, runningSum, 1)
	totalPaths += countPathsWithSum(node.left, targetSum, runningSum, pathCount)
	totalPaths += countPathsWithSum(node.right, targetSum, runningSum, pathCount)
	incrementHashTable(pathCount, runningSum, -1)

	return totalPaths
}

func incrementHashTable(hashTable map[int]int, key int, delta int) {
	newCount := hashTable[key] + delta
	if newCount == 0 {
		delete(hashTable, key)
	} else {
		hashTable[key] = newCount
	}
}

package binarySearch

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	// start at top-right element
	var value int
	row := 0
	col := len(matrix[0]) - 1
	for row <= len(matrix)-1 && col >= 0 {
		value = matrix[row][col]
		if target > value {
			row++
		} else if target < value {
			col--
		} else {
			return true
		}
	}
	return false
}

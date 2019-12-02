package searchMatrix

// divide and conquer, recursion
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return search(matrix, target, 0, len(matrix)-1, 0, len(matrix[0])-1)
}

func search(matrix [][]int, target, r1, r2, c1, c2 int) bool {
	if r1 > r2 || c1 > c2 || r2 >= len(matrix) || c2 >= len(matrix[0]) || c1 < 0 || r1 < 0 {
		return false
	}
	midR := r1 + (r2-r1)/2
	midC := c1 + (c2-c1)/2
	if matrix[midR][midC] == target {
		return true
	} else if matrix[midR][midC] > target {
		a := search(matrix, target, r1, midR-1, c1, midC-1)
		b := search(matrix, target, r1, midR-1, c1, c2)
		c := search(matrix, target, r1, r2, c1, midC-1)
		return a || b || c
	} else {
		a := search(matrix, target, midR+1, r2, midC+1, c2)
		b := search(matrix, target, midR+1, r2, c1, c2)
		c := search(matrix, target, r1, r2, midC+1, c2)
		return a || b || c
	}
}

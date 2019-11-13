package flipAndInvertImage

func flipAndInvertImage(A [][]int) [][]int {
	length := len(A)
	for _, row := range A {
		for i := 0; i*2 < length; i++ {
			if row[i] == row[length-1-i] {
				row[i] ^= 1
				row[length-1-i] = row[i]
			}
		}
	}
	return A
}

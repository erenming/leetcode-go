package MoveZeroes

func moveZeroes(nums []int) {
	index := 0
	for _, i := range nums {
		if i != 0 {
			nums[index] = i
			index++
		}
	}
	for ; index < len(nums); index++ {
		nums[index] = 0
	}
}

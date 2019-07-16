package sorts

func sortColors(nums []int) {
	hi := len(nums) - 1
	gt := hi
	lt := 0
	i := 0
	for i <= gt {
		switch nums[i] {
		case 0:
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		case 1:
			i++
		case 2:
			nums[gt], nums[i] = nums[i], nums[gt]
			gt--
		}
	}
}

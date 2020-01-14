package singleNumber

// 异或定律
// a^0 = a
// a^a = 0
// a^b^a = (a^a)^b
func singleNumber(nums []int) int {
	n := len(nums)
	res := nums[0]
	if n > 1 {
		for i:=1; i<n;i++ {
			res = res ^ nums[i]
		}
	}
	return res
}

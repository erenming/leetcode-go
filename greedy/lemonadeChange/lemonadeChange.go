package lemonadeChange

func lemonadeChange(bills []int) bool {
	var (
		five = 0
		ten  = 0
	)
	for _, value := range bills {
		switch value {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if ten == 0 && five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

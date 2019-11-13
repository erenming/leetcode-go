package majorityElement

import "testing"

func TestMajorityElement(t *testing.T) {
	if majorityElement([]int{6,5,5}) != 5 {
		t.Fail()
	}

}

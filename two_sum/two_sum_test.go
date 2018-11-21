package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expect := []int{0, 1}
	ret := twoSum(nums, target)
	if ret[0] == expect[0] && ret[1] == expect[1] || ret[0] == expect[1] && ret[1] == expect[0] {
		t.Log("succes")
	} else {
		t.Fatal("failed")
	}
}

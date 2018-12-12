package three_sum

/***********************************************************
	https://leetcode-cn.com/problems/3sum/
	给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

	注意：答案中不可以包含重复的三元组。

	例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

	满足要求的三元组集合为：
	[
	[-1, 0, 1],
	[-1, -1, 2]
	]
***********************************************************/

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	var ret [][]int
	sort.Ints(nums)
	end := len(nums) - 1

	for i := 0; i <= end && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := end

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ele := []int{nums[i], nums[left], nums[right]}
				ret = append(ret, ele)

				left++
				for left < right && nums[left-1] == nums[left] {
					left++
				}

				right--
				for left < right && nums[right+1] == nums[right] {
					right--
				}

			} else if sum > 0 {
				right--
				for left < right && nums[right+1] == nums[right] {
					right--
				}
			} else {
				left++
				for left < right && nums[left-1] == nums[left] {
					left++
				}
			}
		}

	}

	return ret
}

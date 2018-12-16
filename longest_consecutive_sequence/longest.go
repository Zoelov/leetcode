package longest_consecutive_sequence

/**********************************************************
	https://leetcode-cn.com/problems/longest-consecutive-sequence/
	定一个未排序的整数数组，找出最长连续序列的长度。

	要求算法的时间复杂度为 O(n)。

	示例:

	输入: [100, 4, 200, 1, 3, 2]
	输出: 4
	解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
**********************************************************/

func longestConsecutive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return 1
	}

	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	max := 0

	for k := range m {
		current := 0
		current++
		delete(m, k)
		positive := 1
		negative := 1
		pos := false
		neg := false
		for {
			if _, ok := m[k+positive]; ok {
				delete(m, k+positive)
				positive++
			} else {
				pos = true
			}
			if _, ok := m[k-negative]; ok {
				delete(m, k-negative)
				negative++
			} else {
				neg = true
			}

			if pos && neg {
				break
			}
		}
		current += (positive + negative - 2)
		if current > max {
			max = current
		}
	}

	return max
}

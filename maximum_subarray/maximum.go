package maximum_subarray

/*********************************************************************************
	https://leetcode-cn.com/problems/maximum-subarray/
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	示例:

	输入: [-2,1,-3,4,-1,2,1,-5,4],
	输出: 6
	解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

************************************************************************************/

func MaxSubArray(nums []int) int {
	ret := -1 << 31
	current := 0
	for i := 0; i < len(nums); i++ {
		current += nums[i]
		if current > ret {
			ret = current
		}

		if current < 0 {
			current = 0
		}
	}

	return ret
}
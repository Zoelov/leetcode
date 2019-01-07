package remove_element

//RemoveElement 改变了元素相对位置
func RemoveElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

// RemoveElement2 修改了原数组，稍微慢一些, 通过append删除指定元素
func RemoveElement2(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

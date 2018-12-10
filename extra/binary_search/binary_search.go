package binary_search

func BinarySearch(data []int, start, end, target int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if data[mid] == target {
		return mid
	}

	if data[mid] > target {
		return BinarySearch(data, start, mid-1, target)
	}

	return BinarySearch(data, mid+1, end, target)
}

func BinarySearch2(data []int, target int) int {
	start := 0
	end := len(data) - 1

	for start <= end {
		mid := (start + end) / 2
		if data[mid] == target {
			return mid
		}

		if data[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

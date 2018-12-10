package sort

func BucketSort(data []int, max int) {
	bucket := make([]int, max)

	for _, v := range data {
		bucket[v]++
	}

	for i, j := 0, 0; i < max; i++ {
		for bucket[i] > 0 {
			data[j] = i
			j++
			bucket[i]--
		}
	}

}

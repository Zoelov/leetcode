package sort

import (
	"leetcode/extra/stack"
)

func QuickSort(data []int, begin, end int) {
	if begin >= end {
		return
	}

	mid := Partition(data, begin, end)

	QuickSort(data, begin, mid-1)
	QuickSort(data, mid+1, end)
}

func Partition(data []int, begin, end int) int {
	baseNum := data[begin]
	i := begin
	j := end

	for i != j {
		for data[j] >= baseNum && i < j {
			j--
		}

		for data[i] <= baseNum && i < j {
			i++
		}

		if i < j {
			data[j], data[i] = data[i], data[j]
		}
	}

	data[begin], data[i] = data[i], data[begin]
	return i
}

func QuickSort2(data []int, begin, end int) {
	s := stack.Stack{}

	if begin < end {
		mid := Partition(data, begin, end)
		if mid-1 > begin {
			s.Push(begin)
			s.Push(mid - 1)
		}

		if mid+1 < end {
			s.Push(mid + 1)
			s.Push(end)
		}

		for !s.IsEmpty() {
			right, _ := s.Pop()
			left, _ := s.Pop()

			p := Partition(data, left.(int), right.(int))
			if p-1 > left.(int) {
				s.Push(left)
				s.Push(p - 1)
			}

			if p+1 < right.(int) {
				s.Push(p + 1)
				s.Push(right)
			}
		}
	}

}

package extra

import "fmt"

/***********************************************
	问题描述：
	写一个栈，并且以O(1)的时间复杂度，来取得其中最小的值

************************************************/

type MiniStack struct {
	data   []int
	length int
}

func (s *MiniStack) Push(val int) {
	element := make([]int, 2)
	if s.length == 0 {
		element[0] = val
		element[1] = val
	} else {
		topMini := s.data[s.length*2-1]

		var mini int
		if topMini > val {
			mini = val
		} else {
			mini = topMini
		}

		element[0] = val
		element[1] = mini

	}

	s.data = append(s.data, element...)
	s.length++
}

func (s *MiniStack) Pop() (int, error) {
	if s.length == 0 {
		return 0, fmt.Errorf("%v", "empty")
	}

	top := s.data[s.length*2-2]
	s.data = s.data[0 : s.length*2-2]
	s.length--

	return top, nil
}

func (s MiniStack) String() string {
	var ret string
	for i := s.length*2 - 1; i >= 0; i-- {
		ret += fmt.Sprintf("%v ", s.data[i])
	}

	return ret
}

func (s *MiniStack) GetMini() (int, error) {
	if s.length == 0 {
		return 0, fmt.Errorf("%v", "empty")
	}

	return s.data[s.length*2-1], nil
}

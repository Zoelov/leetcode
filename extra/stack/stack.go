package stack

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s Stack) Top() interface{} {
	return s[len(s)-1]
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

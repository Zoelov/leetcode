package stack

import "errors"

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

func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("out of index")
	}

	return s[len(s)-1], nil
}

func (s *Stack) Pop() (interface{}, error) {
	theStack := *s
	if theStack.IsEmpty() {
		return nil, errors.New("out of index")
	}

	value := theStack[len(theStack)-1]

	*s = theStack[:len(theStack)-1]

	return value, nil
}

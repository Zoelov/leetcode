package stack

import (
	"testing"
)

func TestMiniStatck(t *testing.T) {
	element := []int{6, 1, 9, 10, 2, -2, 67}

	mini := -2

	s := new(MiniStack)
	for _, v := range element {
		s.Push(v)
	}

	// fmt.Println(s)

	actual, err := s.GetMini()
	if err == nil && actual == mini {
		t.Log("success")
	} else {
		t.Fatal("failed")
	}

	// pop 67
	p, err := s.Pop()
	if p == element[len(element)-1] {
		t.Log("success")
	} else {
		t.Fatal("failed")
	}
	actual, err = s.GetMini()
	if err == nil && actual == mini {
		t.Log("success")
	} else {
		t.Fatal("failed")
	}

	// pop -2
	p, err = s.Pop()
	if p == element[len(element)-2] {
		t.Log("success")
	} else {
		t.Fatal("failed")
	}

	mini = 1
	actual, err = s.GetMini()
	if err == nil && actual == mini {
		t.Log("success")
	} else {
		t.Fatal("failed")
	}

	s.Push(-20)
	mini = -20
	actual, err = s.GetMini()
	if err == nil && actual == mini {
		t.Log("success")
	} else {
		t.Fatal("failed")
	}

}

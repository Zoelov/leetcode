package reverse

import "testing"

func TestReverseInteger(t *testing.T) {
	input := []int{123, -123, 120, 1 << 31}
	expect := []int{321, -321, 21, 0}

	for i := 0; i < len(input); i++ {
		actual := ReverseInteger(input[i])
		if actual != expect[i] {
			t.Fatal("failed")
		}
	}

	t.Log("success")
}

func TestReverseString(t *testing.T) {
	input := "abcdefg"
	expect := "gfedcba"

	actual := ReverseString(input)

	if actual == expect {
		t.Log("success")
	} else {
		t.Fatal("failed")
	}
}

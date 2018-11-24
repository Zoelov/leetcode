package substring_longest

import "testing"

func TestSubstringLongest(t *testing.T) {
	s := "abcabcbb"
	expect := 3

	actual := lengthOfLongestSubstring(s)
	if expect != actual {
		t.Fatal("failed")
	} else {
		t.Log("success")
	}

	s = "bbbbb"
	expect = 1
	actual = lengthOfLongestSubstring(s)
	if expect != actual {
		t.Fatal("failed")
	} else {
		t.Log("success")
	}

	s = "pwwkew"
	expect = 3
	actual = lengthOfLongestSubstring(s)
	if expect != actual {
		t.Fatal("failed")
	} else {
		t.Log("success")
	}
}

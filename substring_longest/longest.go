package substring_longest

/********************************************************

	https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

	示例 1:

	输入: "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
	示例 2:

	输入: "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
	示例 3:

	输入: "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

********************************************************/

func lengthOfLongestSubstring(s string) int {

	length := len(s)
	var ret int
	m := make(map[byte]int)

	for i, j := 0, 0; j < length; j++ {
		if _, ok := m[s[j]]; ok {
			index := m[s[j]]
			if index > i {
				i = index
			}
		}

		if (j - i + 1) > ret {
			ret = j - i + 1
		}
		m[s[j]] = j + 1
	}

	return ret
}

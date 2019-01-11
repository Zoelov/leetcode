package valid_palindrome

/****************************************************************
	https://leetcode-cn.com/problems/valid-palindrome/
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

	说明：本题中，我们将空字符串定义为有效的回文串。

	示例 1:

	输入: "A man, a plan, a canal: Panama"
	输出: true
	示例 2:

	输入: "race a car"
	输出: false
****************************************************************/

func isPalindrome(s string) bool {
	begin, end := 0, len(s)-1

	for begin < end {
		for begin < end && !isCharacterOrNum(s[begin]) {
			begin++
		}

		for begin < end && !isCharacterOrNum(s[end]) {
			end--
		}

		if begin <= end && getUpper(s[begin]) == getUpper(s[end]) {
			begin++
			end--
			continue
		}

		return false
	}

	return true
}

func getUpper(c byte) byte {
	if c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
		return c
	}

	if c >= 'a' && c <= 'z' {
		return c - 32
	}

	return 0
}

func isCharacterOrNum(c byte) bool {
	if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '0' && c <= '9' {
		return true
	}

	return false
}

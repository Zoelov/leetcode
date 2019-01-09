package length_of_last_word

/********************************************************
	https://leetcode-cn.com/problems/length-of-last-word/
	给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。

	如果不存在最后一个单词，请返回 0 。

	说明：一个单词是指由字母组成，但不包含任何空格的字符串。

	示例:

	输入: "Hello World"
	输出: 5
********************************************************/
func LengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	begin := -1
	end := -1
	var hasCharacter bool
	i, j := 0, len(s)-1

	for i <= j {
		for j >= 0 && s[j] == ' ' {
			j--
		}

		if j >= 0 && s[j] != ' ' {
			end = j
			hasCharacter = true
		}

		if s[i] == ' ' && i <= j && s[i+1] != ' ' {
			begin = i
			hasCharacter = true
		}
		i++
	}

	if begin == -1 && end == -1 {
		if hasCharacter {
			return len(s)
		}
		return 0
	}

	if begin == -1 && end != -1 {
		return end + 1
	}

	if begin != -1 && end == -1 {
		return len(s) - begin - 1
	}

	return end - begin
}

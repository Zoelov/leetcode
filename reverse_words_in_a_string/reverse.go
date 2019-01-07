package reverse_words_in_a_string

/*********************************************************
	https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
	给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

	示例 1:

	输入: "Let's take LeetCode contest"
	输出: "s'teL ekat edoCteeL tsetnoc"
	注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
*********************************************************/

func ReverseWords(s string) string {
	s = reverse(s)
	var ret string

	begin := len(s) - 1
	end := len(s)
	for begin >= 0 {
		if s[begin] != ' ' && begin != 0 {
			begin--
			continue
		}

		if begin == 0 {
			ret += s[begin:end]
		} else {
			ret += s[begin+1 : end]
			ret += " "
			end = begin
		}

		begin--
	}

	return ret
}

func reverse(s string) string {
	r := []rune(s)
	begin := 0
	end := len(r) - 1

	for begin < end {
		r[begin], r[end] = r[end], r[begin]
		begin++
		end--
	}

	return string(r)
}

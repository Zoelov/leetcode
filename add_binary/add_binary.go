package add_binary

/****************************************************
	https://leetcode-cn.com/problems/add-binary/
	给定两个二进制字符串，返回他们的和（用二进制表示）。

	输入为非空字符串且只包含数字 1 和 0。

	示例 1:

	输入: a = "11", b = "1"
	输出: "100"
	示例 2:

	输入: a = "1010", b = "1011"
	输出: "10101"
*******************************************************/
func AddBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)

	var ret string
	var add bool
	i, j := aLen-1, bLen-1
	for i >= 0 && j >= 0 {
		if add {
			if a[i] == '1' && b[j] == '1' {
				ret += "1"
				add = true
			} else if a[i] == '1' && b[j] == '0' || a[i] == '0' && b[j] == '1' {
				ret += "0"
				add = true
			} else {
				ret += "1"
				add = false
			}
		} else {
			if a[i] == '1' && b[j] == '1' {
				ret += "0"
				add = true
			} else if a[i] == '1' && b[j] == '0' || a[i] == '0' && b[j] == '1' {
				ret += "1"
				add = false
			} else {
				ret += "0"
				add = false
			}
		}

		i--
		j--
	}

	for j >= 0 {
		if add && b[j] == '1' {
			ret += "0"
			add = true
		} else if add && b[j] == '0' {
			ret += "1"
			add = false
		} else {
			ret += string(b[j])
		}
		j--
	}

	for i >= 0 {
		if add && a[i] == '1' {
			ret += "0"
			add = true
		} else if add && a[i] == '0' {
			ret += "1"
			add = false
		} else {
			ret += string(a[i])
		}
		i--
	}

	if add {
		ret += "1"
	}

	return reverse(ret)

}

func reverse(s string) string {
	r := []rune(s)

	i := 0
	j := len(r) - 1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	return string(r)
}

package zigzag_conversion

/***************************************************
	https://leetcode-cn.com/problems/zigzag-conversion/
	将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

	比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

	L   C   I   R
	E T O E S I I G
	E   D   H   N
	之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

	请你实现这个将字符串进行指定行数变换的函数：

	string convert(string s, int numRows);
	示例 1:

	输入: s = "LEETCODEISHIRING", numRows = 3
	输出: "LCIRETOESIIGEDHN"
	示例 2:

	输入: s = "LEETCODEISHIRING", numRows = 4
	输出: "LDREOEIIECIHNTSG"
	解释:

	L     D     R
	E   O E   I I
	E C   I H   N
	T     S     G
***************************************************/

func Convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	length := len(s)
	b := []byte(s)
	var ret string
	base := 2*numRows - 2
	for i := 1; i <= numRows; i++ {
		if i == 1 || i == numRows {
			j := 0
			for j*base+i-1 < length {
				ret += string(b[j*base+i-1])
				j++
			}

		} else {

			first := base - (2*i - 2)
			second := 2*i - 2
			j := 0
			index := i - 1
			for index < length {
				ret += string(b[index])
				j++
				if j%2 == 1 {
					index += first
				} else {
					index += second
				}
			}
		}
	}

	return ret
}

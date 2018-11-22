package reverse

/***********************************************************************

https://leetcode-cn.com/problems/reverse-integer/

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。
请根据这个假设，如果反转后整数溢出那么就返回 0。

**************************************************************************/

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

func ReverseInteger(x int) int {
	if x > MaxInt32 || x < MinInt32 {
		return 0
	}

	var ret int

	for x != 0 {
		m := x % 10
		x /= 10

		ret = ret*10 + m
	}

	return ret
}

// ReverseString - 顺便写个反转字符串
func ReverseString(s string) string {
	runes := []rune(s)

	i := 0
	j := len(runes) - 1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}

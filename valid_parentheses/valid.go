package valid_parentheses

/*******************************************************************
	https://leetcode-cn.com/problems/valid-parentheses/
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

	有效字符串需满足：

	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	注意空字符串可被认为是有效字符串。

	示例 1:

	输入: "()"
	输出: true
	示例 2:

	输入: "()[]{}"
	输出: true
	示例 3:

	输入: "(]"
	输出: false
	示例 4:

	输入: "([)]"
	输出: false
	示例 5:

	输入: "{[]}"
	输出: true
*********************************************************************/

import "leetcode/extra/stack"

func IsValid(s string) bool {
	p := stack.Stack{}

	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '(' || bytes[i] == '[' || bytes[i] == '{' {
			p.Push(bytes[i])
		}

		if bytes[i] == ')' || bytes[i] == ']' || bytes[i] == '}' {
			if p.IsEmpty() {
				return false
			}

			v := p.Top()
			if bytes[i] == ')' && '(' == v.(byte) || bytes[i] == ']' && '[' == v.(byte) || bytes[i] == '}' && '{' == v.(byte) {
				p.Pop()
			} else {
				return false
			}
		}
	}

	if !p.IsEmpty() {
		return false
	}

	return true
}

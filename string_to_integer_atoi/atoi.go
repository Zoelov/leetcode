package string_to_integer_atoi

import (
	"fmt"
	"strconv"
)

const (
	MaxInt32  = 1<<31 - 1
	MiniInt32 = -1 << 31
)

func MyAtoi(str string) int {
	byteStr := []byte(str)
	if len(byteStr) == 0 {
		return 0
	}

	var positive, negative bool
	var validStr string

	for i := 0; i < len(byteStr); i++ {
		if isNumber(byteStr[i]) {

		}
	}

	for i := 0; i < len(byteStr); i++ {
		if !isNumber(byteStr[i]) {
			if byteStr[i] == '-' {
				if positive {
					break
				}
				negative = true
			} else if byteStr[i] == '+' {
				if negative {
					break
				}
				positive = true
			} else if byteStr[i] == ' ' {
				if len(validStr) == 0 {
					continue
				} else {
					break
				}
			} else {
				break
			}
		} else {
			if len(validStr) == 0 && byteStr[i] == '0' && !positive && !negative {
				continue
			}

			if len(validStr) == 0 && byteStr[i] == '0' && (positive || negative) {
				continue
			}
			validStr += string(byteStr[i])
		}

	}

	if len(validStr) == 0 {
		return 0
	}

	if !positive && !negative {
		positive = true
	}

	if positive {
		compareStr := fmt.Sprintf("%v", MaxInt32)
		if len(validStr) == len(compareStr) && validStr > compareStr || len(validStr) > len(compareStr) {
			return MaxInt32
		}
	}

	if negative {
		s := "-" + validStr
		validStr = s
		compareStr := fmt.Sprintf("%v", MiniInt32)
		if len(s) == len(compareStr) && s > compareStr || len(s) > len(compareStr) {
			return MiniInt32
		}
	}

	v, _ := strconv.Atoi(validStr)
	return v
}

func isNumber(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}

	return false
}

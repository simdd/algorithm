package problem

import (
	"strings"
)

func longestValidParenthesesByCount(s string) int {
	origin := strings.Split(s, "")

	maxlenleft := 0
	maxlenright := 0

	left := 0
	right := 0

	for _, v := range origin {
		if v == "(" {
			left++
		} else if v == ")" {
			right++
		}

		if right > left {
			right = 0
			left = 0
		} else if right == left {
			if maxlenleft < left+right {
				maxlenleft = right + left
			}
		}
	}

	left = 0
	right = 0

	for i := len(origin) - 1; i >= 0; i-- {
		v := origin[i]

		if v == "(" {
			left++
		} else if v == ")" {
			right++
		}

		if left > right {
			right = 0
			left = 0
		} else if left == right {
			if maxlenright < right+left {
				maxlenright = left + right
			}
		}
	}

	if maxlenleft < maxlenright {
		return maxlenright
	}

	return maxlenleft
}

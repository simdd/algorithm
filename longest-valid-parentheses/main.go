package problem

import (
	"strings"
)

func longestValidParenthesesByCount(s string) []string {
	origin := strings.Split(s, "")
	maxlen := 0
	left := 0
	right := 0

	for _, v := range origin {
		if v == "(" {
			left++
		} else if v == ")" {
			right++
		}

		if right > left {
			if maxlen < right+left {
				maxlen = right + left
			}

			right = 0
			left = 0
		}
	}

	return origin
}

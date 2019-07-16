package problem

import (
	"strings"
)

func byCount(s string) int {
	origin := strings.Split(s, "")

	var maxlenleft int
	var maxlenright int

	var left int
	var right int

	for _, v := range origin {
		if v == "(" {
			left++
		} else {
			right++
		}

		if right > left {
			right = 0
			left = 0
		}

		if right == left {
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
		} else {
			right++
		}

		if left > right {
			right = 0
			left = 0
		}

		if left == right {
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

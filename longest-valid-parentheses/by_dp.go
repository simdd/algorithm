package problem

import (
	"strings"
)

func byDp(s string) int {
	o := strings.Split(s, "")
	dp := make([]int, len(o))
	max := 0

	for i := 1; i < len(o); i++ {
		if o[i] == ")" {
			if o[i-1] == "(" {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = +2
				}
			} else if i-dp[i-1] > 0 && o[i-dp[i-1]-1] == "(" {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

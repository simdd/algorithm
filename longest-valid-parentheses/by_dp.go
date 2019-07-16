package problem

import (
	"strings"
)

func byDp(s string) int {
	o := strings.Split(s, "")
	dp := make([]int, 1)

	for i := 1; i < len(o); i++ {
		if o[i] == "(" {
			dp = append(dp, dp[i-1])
		} else {
			if o[i-1] == "(" {
				if i-2 < 0 {
					dp = append(dp, 2)
				} else {
					dp = append(dp, dp[i-2]+2)
				}
			} else {
				if i-dp[i-1] > 0 && o[i-dp[i-1]-1] == "(" {
					if i-dp[i-1]-2 > 0 {
						dp = append(dp, dp[i-1]+2+dp[i-dp[i-1]-2])
					} else {
						dp = append(dp, dp[i-1]+2)
					}
				} else {
					dp = append(dp, dp[i-1])
				}
			}
		}
	}

	return dp[len(dp)-1]
}

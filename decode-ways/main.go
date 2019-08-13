package problem

import (
	"strconv"
	"strings"
)

func main(s string) int {
	strs := strings.Split(s, "")
	o := make([]int, len(strs))
	for i := range o {
		o[i], _ = strconv.Atoi(strs[i])
	}

	dp := make([]int, len(o))

	if o[0] == 0 {
		dp[0] = 0
		return dp[0]
	}

	dp[0] = 1

	if len(o) == 1 {
		return dp[0]
	}

	if o[1] == 0 && o[0] > 2 {
		return 0
	} else if o[1] == 0 && o[0] <= 2 {
		dp[1] = 1
	} else if o[0]*10+o[1] > 26 {
		dp[1] = 1
	} else {
		dp[1] = 2
	}

	if len(o) == 2 {
		return dp[1]
	}

	// done
	for i := 2; i < len(o); i++ {
		if o[i] == 0 {
			if o[i-1] == 1 || o[i-1] == 2 {
				return dp[i-2]
			}

			return 0
		}

		if o[i-1] == 0 || o[i-1] > 6 {
			return dp[i-1]
		} else if o[i-1]*10+o[i] > 26 {
			return dp[i-1]
		} else {
			return dp[i-1] + dp[i-2]
		}
	}

	return dp[len(dp)-1]
}

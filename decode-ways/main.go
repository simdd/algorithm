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

	if o[0] != 0 {
		dp[0] = 1
	} else {
		dp[0] = 0
	}

	if len(o) == 1 {
		return dp[0]
	}

	if o[0] > 2 || o[1] > 6 {
		dp[1] = 1
	} else if o[0] == 0 {
		dp[1] = 0
	} else if o[1] == 0 {
		dp[1] = 1
	} else {
		dp[1] = 2
	}

	if len(o) == 2 {
		return dp[1]
	}

	for i := 2; i < len(o); i++ {
		if o[i] == 0 && (o[i-1] == 0 || o[i-1] > 2) {
			return 0
		}

		if o[i-1] > 0 && o[i-1] < 3 && o[i] > 0 && o[i] < 7 {
			dp[i] = dp[i-1] + dp[i-2]
		} else if o[i] == 0 {
			dp[i] = dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(dp)-1]
}

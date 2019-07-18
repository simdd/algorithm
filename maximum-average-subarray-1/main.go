package problem

import (
	"math"
)

func main(nums []int, k int) float64 {
	max := float64(-10000)

	for i := k - 1; i < len(nums); i++ {
		sum := 0

		for n := i - (k - 1); n <= i; n++ {
			sum = sum + nums[n]
		}

		var cur = float64(sum) / float64(k)

		max = math.Max(cur, max)
	}

	return max
}

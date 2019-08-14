package problem

import "math"

func main(x int) (out int) {
	for ; x != 0; x /= 10 {
		out = out*10 + x%10

		if out > math.MaxInt32 || out < -math.MaxInt32-1 {
			return 0
		}
	}
	return out
}

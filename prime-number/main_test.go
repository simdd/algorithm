package problem

import "testing"
import "fmt"

func StringSliceEqualBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestPrimeNumber(t *testing.T) {
	origin := 40

	target := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
	}

	res := StringSliceEqualBCE(primeNumber(origin), target)

	fmt.Printf("%v\n", res)
	fmt.Printf("%v", primeNumber(origin))
}

package problem

import "testing"

func TestCountPrimes(t *testing.T) {
	origin := 40
	target := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
	}

	ret := countPrimes(origin)

	if ret != len(target) {
		t.Error("TestCountPrimes ", ret)
	}
}

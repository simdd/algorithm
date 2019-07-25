package problem

import (
	"algorithm/utils"
	"testing"
)

func TestPrimeNumber(t *testing.T) {
	origin := 40

	target := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
	}

	ret := utils.IntSliceEqualBCE(primeNumber(origin), target)

	if !ret {
		t.Error("TestPrimeNumber: ", ret)
	}
}

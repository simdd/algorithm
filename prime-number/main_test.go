package problem

import (
	"algorithm/utils"
	"fmt"
	"testing"
)

func TestPrimeNumber(t *testing.T) {
	origin := 40

	target := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
	}

	res := utils.IntSliceEqualBCE(primeNumber(origin), target)

	fmt.Printf("%v\n", res)
	fmt.Printf("%v", primeNumber(origin))
}

package problem

import (
	"algorithm/utils"
	"testing"
)

func TestMain(t *testing.T) {
	ret := main([]int{3, 5, 1, 2, 6, 2, 4})
	target := []int{1, 2, 2, 3, 3, 3, 4}

	if !utils.IntSliceEqualBCE(ret, target) {
		t.Error("TestMain ", ret)
	}
}

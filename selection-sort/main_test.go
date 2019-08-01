package problem

import (
	"algorithm/utils"
	"testing"
)

func TestMain(t *testing.T) {
	ret := main([]int{3, 5, 1, 2, 2, 2, 2})
	target := []int{1, 2, 2, 2, 2, 3, 5}

	if !utils.IntSliceEqualBCE(ret, target) {
		t.Error("TestMain ", ret)
	}
}

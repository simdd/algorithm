package problem

import (
	"algorithm/utils"
	"testing"
)

func TestMain(t *testing.T) {
	ret := main([]int{6, 7, 2, 3, 2, 1, 8, 9, 3, 4})
	target := []int{1, 2, 2, 3, 3, 4, 6, 7, 8, 9}

	if !utils.IntSliceEqualBCE(ret, target) {
		t.Error("TestMain ", ret)
	}
}

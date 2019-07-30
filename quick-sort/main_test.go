package problem

import (
	"algorithm/utils"
	"testing"
)

func TestMain(t *testing.T) {
	ret := main([]int{1, 3, 2, 3, 3, 2, 4})
	target := []int{1, 2, 2, 3, 3, 3, 4}

	if !utils.IntSliceEqualBCE(ret, target) {
		t.Error("TestMain ", ret)
	}
}

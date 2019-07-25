package problem

import (
	"testing"
)

func TestMain(t *testing.T) {
	ret := main([]int{-1, 1, 2, 3}, 2)

	if ret != 2.5 {
		t.Error("TestMain ", ret)
	}
}

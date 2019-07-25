package problem

import (
	"testing"
)

func TestMain(t *testing.T) {
	ret := main(4)
	target := 5

	if ret != target {
		t.Error("TestMain ", ret)
	}
}

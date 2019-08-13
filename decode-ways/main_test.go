package problem

import (
	"testing"
)

func TestMain(t *testing.T) {
	ret := main("110")

	if ret != 1 {
		t.Error("Decode Ways ", ret)
	}
}

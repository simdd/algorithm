package problem

import (
	"testing"
)

func TestByDp(t *testing.T) {
	ret := byDp("(()())")

	if ret != 6 {
		t.Error("TestByDp: ", ret)
	}
}

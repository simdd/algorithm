package problem

import (
	"testing"
)

func TestMain(t *testing.T) {
	ret := main("MCMXCIV")

	if ret != 1994 {
		t.Error("Roman to Integer: ", ret)
	}
}

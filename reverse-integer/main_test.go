package problem

import (
	"testing"
)

func TestMain(t *testing.T) {
	ret := main(1234)

	if ret != 4321 {
		t.Error("reverse integer ", ret)
	}
}

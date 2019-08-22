package problem

import "testing"

func TestMain(t *testing.T) {
	ret := main([]string{})

	if ret != "fl" {
		t.Error("longest common prefix: ", ret)
	}
}

package problem

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	ret := countAndSay(5)
	target := "111221"

	if ret != target {
		t.Error("TestCountAndSay  ", ret)
	}
}

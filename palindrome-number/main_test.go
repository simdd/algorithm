package problem

import (
	"testing"
)

func TestMain(t *testing.T) {
	ret := main(1221)

	if ret != true {
		t.Error("palindrome-number:  ", ret)
	}
}

package problem

import (
	"fmt"
	"testing"
)

func TestLongestValidParenthesesByCount(t *testing.T) {
	ret := longestValidParenthesesByCount("(())))(()")
	fmt.Println(ret)
}

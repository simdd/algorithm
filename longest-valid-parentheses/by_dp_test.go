package problem

import (
	"fmt"
	"testing"
)

func TestByDp(t *testing.T) {
	ret := byDp("(()())")
	fmt.Println(ret)
}

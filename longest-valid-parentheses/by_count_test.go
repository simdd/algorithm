package problem

import (
	"fmt"
	"testing"
)

func TestByCount(t *testing.T) {
	ret := byCount("()((((())))")
	fmt.Println(ret)
}

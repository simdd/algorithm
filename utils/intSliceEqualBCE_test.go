package utils

import (
	"fmt"
	"testing"
)

func TestIntSliceEqualBCE(t *testing.T) {
	ret := intSliceEqualBCE([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(ret)
}

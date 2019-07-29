package problem

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	ret := main([]int{1, 3, 2, 4})
	// target := []int{1, 2, 3, 4}
	fmt.Println(ret)

	// if !utils.IntSliceEqualBCE(ret, target) {
	// 	t.Error("TestMain ", ret)
	// }
}

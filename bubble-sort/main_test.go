package problem

import (
	"algorithm/utils"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	ret := utils.IntSliceEqualBCE(main2([]int{1, 3, 2, 4}), []int{1, 2, 3})
	fmt.Println(ret)
}
